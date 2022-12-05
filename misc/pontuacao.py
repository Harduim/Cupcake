import os
from sqlmodel import SQLModel as _SQLModel, select, create_engine, Field
from sqlmodel.orm.session import Session
from sqlalchemy.future.engine import Engine
from dotenv import load_dotenv
from dataclasses import dataclass
from uuid import UUID

load_dotenv()

DB_USERNAME = os.getenv("DB_USERNAME")
DB_PASSWORD = os.getenv("DB_PASSWORD")
DB_DATABASE = os.getenv("DB_DATABASE")
DB_PORT = os.getenv("DB_PORT")
DB_HOST = os.getenv("DB_HOST")

OITAVAS_ID = UUID("ef13e77f-b345-4f4d-b4a7-2d1cfb12fa48")

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

    GROUP_STAGE = 3


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
    id: UUID = Field(primary_key=True)
    name: str
    national_team_a_id: UUID
    national_team_b_id: UUID
    gol_a: int
    gol_b: int
    bracket_id: UUID
    winner_id: UUID


class NationalTeam(SQLModel, table=True):
    __tablename__ = "national_teams"
    id: UUID = Field(primary_key=True)
    name: str


class Bet(SQLModel, table=True):
    __tablename__ = "bets"
    id: UUID = Field(primary_key=True)
    gol_a: int
    gol_b: int
    user_id: UUID
    match_id: UUID
    winner_id: UUID


class Bracket(SQLModel, table=True):
    __tablename__ = "brackets"
    id: UUID = Field(primary_key=True)
    name: str
    multiplier: int


class GroupStage(SQLModel, table=True):
    __tablename__ = "groups"
    user_id: str = Field(primary_key=True)
    national_team_id: str = Field(primary_key=True)


class User(SQLModel, table=True):
    __tablename__ = "users"
    id: UUID = Field(primary_key=True)
    name: str
    points: int

    def bets(self) -> list[Bet]:
        with Session(conn) as db:
            return db.exec(select(Bet).where(Bet.user_id == self.id)).all()

    def group_stage_bets(self) -> list[GroupStage]:
        with Session(conn) as db:
            return db.exec(select(GroupStage).where(GroupStage.user_id == str(self.id))).all()

    def update_score(self, score: int):
        self.points = score
        with Session(conn) as db:
            db.add(self)
            db.commit()


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
    user: User, teams: dict[str, NationalTeam], brackets: dict[str, Bracket], matches: list[Match]
):
    round_16_matches = [m for m in matches if m.bracket_id == OITAVAS_ID]

    round_16_teams = {
        *{m.national_team_a_id for m in round_16_matches},
        *{m.national_team_b_id for m in round_16_matches},
    }
    group_stage_bets = {UUID(gb.national_team_id) for gb in user.group_stage_bets()}
    acertos = {g for g in group_stage_bets if g in round_16_teams}
    group_stage_points = len(acertos) * Points.GROUP_STAGE

    final_score = group_stage_points
    user_bets = user.bets()

    def format_aposta(nat_teams: set[UUID]) -> str:
        return ", ".join(sorted([teams[gb].name for gb in nat_teams]))

    print("\n" + user.name)
    print("\n\tFase de grupos")
    print(f"\tAposta:    {format_aposta(group_stage_bets)}")
    print(f"\tResultado: {format_aposta(round_16_teams)}")
    print(f"\t{len(acertos)} acertos: {format_aposta(acertos)}")
    print(f"\tPontos da fase de grupos: {group_stage_points}")

    match_map = {m.id: m for m in matches}

    for bet in user_bets:
        match = match_map[bet.match_id]
        if match.winner_id is None:
            # Skipping matches that did ont have happened yet
            continue

        bracket = brackets[match.bracket_id]
        score = calculate_match_score(bet, match, teams, bracket.multiplier)
        print(f"\tPontuação anterior: {final_score}")
        final_score += score
        print(f"\tPontuação atualizada: {final_score}")

    return final_score


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
    matches = Match.all()
    teams: dict[str, NationalTeam] = NationalTeam.id_map()
    brackets: dict[str, Bracket] = Bracket.id_map()
    for user in User.all():
        final_score = calculate_score(user, teams, brackets, matches)
        user.update_score(final_score)
