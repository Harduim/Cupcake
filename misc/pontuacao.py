import os
from sqlmodel import SQLModel as _SQLModel, select, create_engine, Field
from sqlmodel.orm.session import Session
from sqlalchemy.future.engine import Engine
from dotenv import load_dotenv
from dataclasses import dataclass

load_dotenv()

DB_USERNAME = os.getenv("DB_USERNAME")
DB_PASSWORD = os.getenv("DB_PASSWORD")
DB_DATABASE = os.getenv("DB_DATABASE")
DB_PORT = os.getenv("DB_PORT")
DB_HOST = os.getenv("DB_HOST")

assert all([DB_USERNAME, DB_PASSWORD, DB_DATABASE, DB_PORT, DB_HOST]), "Missing environment vars"

conn: Engine | None = None


@dataclass
class Points:
    # Amount of points
    GOAL = 1

    # Cannot infer winner because of ties an penalty shootouts
    SCOREBOARD = 2
    WINNER = 2
    SCOREBOARD_WINNER = 3


class SQLModel(_SQLModel):
    @classmethod
    def all(cls) -> list[_SQLModel]:
        with Session(conn) as db:
            return db.exec(select(cls)).all()

    @classmethod
    def id_map(cls) -> dict[str, _SQLModel]:
        return {m.id: m for m in cls.all()}


class Match(SQLModel, table=True):
    __tablename__ = "matches"
    id: str = Field(primary_key=True)
    name: str
    national_team_a_id: str
    national_team_b_id: str
    gol_a: int
    gol_b: int
    bracket_id: str
    winner_id: str


class NationalTeam(SQLModel, table=True):
    __tablename__ = "national_teams"
    id: str = Field(primary_key=True)
    name: str


class Bet(SQLModel, table=True):
    __tablename__ = "bets"
    id: str = Field(primary_key=True)
    gol_a: int
    gol_b: int
    user_id: str
    match_id: str
    winner_id: str


class User(SQLModel, table=True):
    __tablename__ = "users"
    id: str = Field(primary_key=True)
    name: str
    points: int

    def bets(self) -> list[Bet]:
        with Session(conn) as db:
            return db.exec(select(Bet).where(Bet.user_id == self.id)).all()


class Bracket(SQLModel, table=True):
    __tablename__ = "brackets"
    id: str = Field(primary_key=True)
    name: str
    multiplier: int


class GroupStage(SQLModel, table=True):
    __tablename__ = "groups"
    user_id: str = Field(primary_key=True)
    national_team_id: str = Field(primary_key=True)


def calculate_match_score(
    bet: Bet,
    match: Match,
    teams: dict[str, NationalTeam],
    multiplier: int,
):

    selecao_a = teams[match.national_team_a_id]
    selecao_b = teams[match.national_team_b_id]

    GOAL = Points.GOAL * multiplier
    SCOREBOARD = Points.SCOREBOARD * multiplier
    WINNER = Points.WINNER * multiplier
    SCOREBOARD_WINNER = Points.SCOREBOARD_WINNER * multiplier

    user_scores = [0]

    print(
        f"""
        Aposta\t  - {selecao_a.name}: {bet.gol_a} x {selecao_b.name}: {bet.gol_b}
        Resultado - {selecao_a.name}: {match.gol_a} x {selecao_b.name}: {match.gol_b}
        """
    )

    if bet.gol_a == match.gol_a:
        user_scores.append(GOAL)
        print(f"\tPontos ganhos por acertar os gols seleção {selecao_a.name}: {GOAL} ")

    if bet.gol_b == match.gol_b:
        user_scores.append(GOAL)
        print(f"\tPontos ganhos por acertar os gols seleção {selecao_b.name}: {GOAL} ")

    if bet.gol_a == match.gol_a and bet.gol_b == match.gol_b:
        user_scores.append(SCOREBOARD)
        print(f"\tPontos ganhos por acertar o placar do jogo: {SCOREBOARD} ")

    if bet.winner_id == match.winner_id:
        user_scores.append(WINNER)
        print(f"\tPontos ganhos por acertar a seleção vencedora da partida: {WINNER} ")

    if bet.gol_a == match.gol_a and bet.gol_b == match.gol_b and bet.winner_id == match.winner_id:
        user_scores.append(SCOREBOARD_WINNER)
        print(f"\tPontos ganhos por acertar o placar e a seleção vencedora: {SCOREBOARD_WINNER} ")

    return sum(user_scores)


def calculate_score(
    user: User,
    teams: dict[str, NationalTeam],
    brackets: dict[str, Bracket],
    group_stage_points: int = 0,
):
    final_score = group_stage_points
    user_bets = user.bets()
    if len(user_bets) == 0:
        return 0
    print("\n" + user.name)
    if group_stage_points>0:
        print(f"\tPontos da fase de grupos: {group_stage_points}")
    for bet in user_bets:
        match = matches[bet.match_id]
        if match.winner_id is None:
            # Skipping matches that did ont have happened yet
            continue

        bracket = brackets[match.bracket_id]
        score = calculate_match_score(bet, match, teams, bracket.multiplier)
        print(f"\tPontuação anterior: {final_score}")
        final_score += score
        print(f"\tPontuação atualizada: {final_score}")


if __name__ == "__main__":
    conn = create_engine(
        "postgresql+psycopg2://{}:{}@{}:{}/{}".format(
            DB_USERNAME,
            DB_PASSWORD,
            DB_HOST,
            DB_PORT,
            DB_DATABASE,
        )
    )
    matches: dict[str, Match] = Match.id_map()
    teams: dict[str, NationalTeam] = NationalTeam.id_map()
    brackets: dict[str, Bracket] = Bracket.id_map()
    for user in User.all():
        calculate_score(user, teams, brackets)
