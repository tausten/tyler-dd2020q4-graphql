// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Animal interface {
	IsAnimal()
}

type Named interface {
	IsNamed()
}

type Node interface {
	IsNode()
}

type Zone interface {
	IsZone()
}

type ZoneOccupants interface {
	IsZoneOccupants()
}

type Barn struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Occupants []ZoneOccupants `json:"occupants"`
}

func (Barn) IsNode()  {}
func (Barn) IsNamed() {}
func (Barn) IsZone()  {}

type Chicken struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	AnimalSpecificStuff string `json:"animalSpecificStuff"`
}

func (Chicken) IsNode()          {}
func (Chicken) IsNamed()         {}
func (Chicken) IsAnimal()        {}
func (Chicken) IsZoneOccupants() {}

type Coop struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Occupants []ZoneOccupants `json:"occupants"`
}

func (Coop) IsNode()  {}
func (Coop) IsNamed() {}
func (Coop) IsZone()  {}

type Farm struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (Farm) IsNode()  {}
func (Farm) IsNamed() {}

type FarmHouse struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Occupants []ZoneOccupants `json:"occupants"`
}

func (FarmHouse) IsNode()  {}
func (FarmHouse) IsNamed() {}
func (FarmHouse) IsZone()  {}

type Field struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Occupants []ZoneOccupants `json:"occupants"`
}

func (Field) IsNode()  {}
func (Field) IsNamed() {}
func (Field) IsZone()  {}

type Horse struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	AnimalSpecificStuff string `json:"animalSpecificStuff"`
}

func (Horse) IsNode()          {}
func (Horse) IsNamed()         {}
func (Horse) IsAnimal()        {}
func (Horse) IsZoneOccupants() {}

type NewTodo struct {
	Text   string `json:"text"`
	UserID string `json:"userId"`
}

type Person struct {
	ID                  string `json:"id"`
	Name                string `json:"name"`
	AnimalSpecificStuff string `json:"animalSpecificStuff"`
}

func (Person) IsNode()          {}
func (Person) IsNamed()         {}
func (Person) IsAnimal()        {}
func (Person) IsZoneOccupants() {}

type QueryFarmResult struct {
	Data   *Farm    `json:"data"`
	Errors []string `json:"errors"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
