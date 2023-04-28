import json


def unifica_dados():
    # Ler dados do primeiro arquivo JSON
    with open("./processa_dados/data/Taseif_characters.json", "r") as f:
        data1 = json.load(f)

    # Ler dados do segundo arquivo JSON
    with open("./processa_dados/data/Counterplay_characters.json", "r") as f:
        data2 = json.load(f)

    # Criar um novo dicionário vazio
    data_unificado = []

    # Adicionar os dados de cada arquivo JSON ao dicionário unificado
    for data in data1:
        data_unificado.append(data)
    for data in data2:
        data_unificado.append(data)

    # Escrever o dicionário unificado em um novo arquivo JSON
    with open("arquivo_unificado.json", "w") as f:
        json.dump(data_unificado, f)


if __name__ == "__main__":
    unifica_dados()
