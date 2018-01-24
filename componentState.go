package main

/**
 * Created by John Tsantilis (A.K.A lumi) on 18/1/2018.
 * @author John Tsantilis <i.tsantilis [at] yahoo [dot] com>
 */

//All possible states: "BOOTSTRAPPED", "INITIALIZED", "DEPLOYED", "BLOCKED", "STARTED", "STOPPED", "UNDEPLOYED", "ERRONEOUS", "CHAINED"
type ComponentState struct {
	State string `json:"state"`

}

func NewComponentStateEmpty() *ComponentState {
	return &ComponentState{}

}

func NewComponentState(state string) *ComponentState {
	return &ComponentState{
		State: state,

	}

}

func (c *ComponentState) GetState() (state string) {
	return c.State

}

func (c *ComponentState) SetState(state string) {
	c.State = state

}