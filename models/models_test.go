package models

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorld(t *testing.T) {

	t.Run("Retorno de String", func(t *testing.T) {
		actualString := World{Name: "Inabra"}.String()
		expectedString := "World: Inabra"
		assert.Equal(t, expectedString, actualString)
	})

	t.Run("Teste de Tipagem", func(t *testing.T) {
		w := World{
			Name:      "Inabra",
			Guilds:    []string{"Taseif", "Counterplay"},
			BattleEye: "Yellow",
		}
		actualType := reflect.TypeOf(w)
		expectedType := reflect.TypeOf(World{})

		assert.Equal(t, actualType, expectedType)
	})
}

func TestCharacter(t *testing.T) {
	t.Run("Teste de Tipagem", func(t *testing.T) {
		c := Character{
			CharacterName: "Yukiezera General",
			FormerNames:   []string{"Yukie Mangote Loko"},
			Level:         1306,
		}

		actualType := reflect.TypeOf(c)
		expectedType := reflect.TypeOf(Character{})

		assert.Equal(t, actualType, expectedType)
	})

}

func TestGuild(t *testing.T) {
	g := Guild{
		World{
			Name:      "Inabra",
			Guilds:    []string{"Zenobra"},
			BattleEye: "Green",
		},
		"Zenobra Pune",
		420,
	}
	actualType := reflect.TypeOf(g)
	expectedType := reflect.TypeOf(Guild{})

	assert.Equal(t, actualType, expectedType)
}
