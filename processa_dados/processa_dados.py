import requests
import json
import os

guilds = ["Taseif", "Counterplay"]


def get_guild_members(guild_name: str) -> list:
    url = f"https://dev.tibiadata.com/v4/guild/{guild_name}"
    response = requests.get(url)
    data = response.json()

    members = []
    for member in data.get("guild").get("members"):
        member_name = member.get("name")
        member_level = member.get("level")

        member_info = {
            "CharacterName": member_name,
            "Level": member_level,
            "Guild": guild_name,
        }
        if member.get("level") > 555:
            members.append(member_info)

    return members


def create_guild_members_json(guild_name: str) -> bool:
    members = get_guild_members(guild_name)
    return create_json(members, guild_name)


def create_json(
    characters: list,
    guild_name: str,
) -> bool:
    path = f"data/{guild_name}_characters.json"

    with open(path, "w") as f:
        json.dump(characters, f)

    return os.path.exists(path)


def update_characters_exp(
    characters_path: str,
    world: str,
    index: int,
) -> bool:
    with open(characters_path, "r") as f:
        characters = json.load(f)

    url = f"https://dev.tibiadata.com/v4/highscores/{world}/experience/all/{index}"
    response = requests.get(url=url)
    if response.ok:
        data = response.json()
        highscore_list = data.get("highscores").get("highscore_list")

        for char in characters:
            highscore_char = next(
                (hs for hs in highscore_list if hs["name"] == char["CharacterName"]),
                None,
            )
            if highscore_char:
                char["Exp"] = highscore_char["value"]

            if not char.get("Exp"):
                char["Exp"] = 0

        with open(characters_path, "w") as f:
            json.dump(characters, f)

        return True

    return False


def full_process_alchemist(guild_name, world):
    json_path = f"data/{guild_name}_characters.json"

    if not create_guild_members_json(guild_name):
        return False

    for i in range(21):
        if i < 1:
            continue
        if not update_characters_exp(json_path, world, i):
            return False
    print(f"Dados processados de {guild_name}")
    return True


def unifica_dados():
    # Ler dados do primeiro arquivo JSON
    with open("data/Taseif_characters.json", "r") as f:
        data1 = json.load(f)

    # Ler dados do segundo arquivo JSON
    with open("data/Counterplay_characters.json", "r") as f:
        data2 = json.load(f)

    # Criar um novo dicionário vazio
    data_unificado = []

    # Adicionar os dados de cada arquivo JSON ao dicionário unificado
    for data in data1:
        data_unificado.append(data)
    for data in data2:
        data_unificado.append(data)

    # Escrever o dicionário unificado em um novo arquivo JSON
    with open("data/arquivo_unificado.json", "w") as f:
        json.dump(data_unificado, f)


if __name__ == "__main__":
    world = "Inabra"
    for guild_name in guilds:
        full_process_alchemist(guild_name, world)

    unifica_dados()
