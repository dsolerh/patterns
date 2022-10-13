package main

import (
	"fmt"
	"sync"
)

// CoR, Mediator, Observer, CQS

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(q *Query)
}

type Observable interface {
	Subscribe(o Observer)
	UnSubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
}

func (g *Game) UnSubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value any) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type Creature struct {
	game            *Game
	Name            string
	attack, defense int
}

func NewCreature(g *Game, name string, attack, defense int) *Creature {
	return &Creature{g, name, attack, defense}
}

func (c *Creature) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *Creature) String() string {
	return fmt.Sprintf("%s (%d/%d)", c.Name, c.Attack(), c.Defense())
}

type CreatureModifier struct {
	game     *Game
	creature *Creature
}

func (c *CreatureModifier) Handle(q *Query) {

}

type DoubleAttackMod struct {
	CreatureModifier
}

func (c *DoubleAttackMod) Handle(q *Query) {
	if q.CreatureName == c.creature.Name &&
		q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackMod) Close() error {
	d.game.UnSubscribe(d)
	return nil
}

func NewDoubleAttackMod(g *Game, c *Creature) *DoubleAttackMod {
	d := &DoubleAttackMod{CreatureModifier{g, c}}
	g.Subscribe(d)
	return d
}

func main() {
	game := &Game{sync.Map{}}
	goblin := NewCreature(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.String())

	{
		m := NewDoubleAttackMod(game, goblin)
		fmt.Println(goblin.String())
		m.Close()
	}

	fmt.Println(goblin.String())
}
