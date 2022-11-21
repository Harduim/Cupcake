import pandas as pd
from dataclasses import make_dataclass
import itertools

#multiplier - column 'multiplier' from brackets
key_title = ['Final','Disputa pelo 3','Semi Final','Quartas','Oitavas','Rodada Coringa']
multiplier_value =  [8,3,3,2,1,9]
multiplier_keys = dict(zip(key_title, multiplier_value))

#bets
Bet = make_dataclass("Bet", [("id", int), ("timestamp", str), ("selecao_a", str), ("selecao_b", str)
                            , ("gols_a", int), ("gols_b", int), ("user_id", int), ("partida_id", int), ("winner_id", int)])

Match = make_dataclass("Match", [("id", int), ("timestamp", str), ("selecao_a", str), ("selecao_b", str)
                            , ("gols_a", int), ("gols_b", int), ("bracket_id", str), ("winner_id", int)])

bet1 = Bet(1,"2022-12-03","Brasil","Argentina",1,1,7,10,3)
bet5 = Bet(5,"2022-12-09","Chile","Uruguai",3,3,9,11,4)

match1 = Match(18,"2022-12-09","Brasil","Argentina",1,1,"Oitavas",10,3)
match2 = Match(19,"2022-12-09","Chile","Uruguai",0,0,"Rodada Coringa",11,3)

def calculate_bet_score(bet,match):
    #scores
    goals = 1
    scoreboard = 2
    winner = 2
    full_score = 3 

    user_scores = []

    print(f'''{bet.user_id} vs Resultados
              User {bet.user_id} - {bet.selecao_a}: {bet.gols_a} x {bet.selecao_b}: {bet.gols_b}
              Resultado oficial - {match.selecao_a}: {match.gols_a} x {match.selecao_b}: {match.gols_b}''')

    if bet.gols_a == match.gols_a:
        user_scores.append(goals)
        print(f"Pontos ganhos por acertar os gols seleção {bet.selecao_a}: {goals} ")

    if bet.gols_b == match.gols_b:
        user_scores.append(goals)
        print(f"Pontos ganhos por acertar os gols seleção {bet.selecao_b}: {goals} ")

    if bet.gols_a == match.gols_a and bet.gols_b == match.gols_b:
        user_scores.append(scoreboard)
        print(f"Pontos ganhos por acertar o placar do jogo: {scoreboard} ")

    if bet.winner_id == match.winner_id:
        user_scores.append(winner)
        print(f"Pontos ganhos por acertar a seleção vencedora da partida: {winner} ")
    
    if bet.gols_a == match.gols_a and bet.gols_b == match.gols_b and bet.winner_id == match.winner_id:
        user_scores.append(full_score)
        print(f"Pontos ganhos por acertar o placar e a seleção vencedora: {full_score} ")

    scores_total = sum(user_scores)
    print(f"Pontuação total sem fator multiplicador: {scores_total}")

    multiplier = multiplier_keys[match.bracket_id]
    print(f"Multiplicador da chave {match.bracket_id}: {multiplier}")

    final_score = scores_total*multiplier
    print(f"Pontuação total com multiplicador: {final_score}")

    return final_score

users = ["Default"]
matches = [match1,match2]
bets = [bet1,bet5]

for user in users: 
    final_score = 0
    # a aposta (bet) vai ser pega com o id da jogo e o id do usuário 
    for bet,match in zip(bets,matches):
        score = calculate_bet_score(bet,match)
        (print(f"Pontuação anterior: {final_score}"))
        final_score += score
        (print(f"Pontuação atualizada: {final_score}"))




    


