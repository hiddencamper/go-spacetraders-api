package APITools

import "time"

type GetStatus struct {
	Status        string         `json:"status"`
	Version       string         `json:"version"`
	ResetDate     string         `json:"resetDate"`
	Description   string         `json:"description"`
	Stats         Stats          `json:"stats"`
	Leaderboards  Leaderboards   `json:"leaderboards"`
	ServerResets  ServerResets   `json:"serverResets"`
	Announcements []Announcement `json:"announcements"`
	Links         []Link         `json:"links"`
}

type Announcement struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Leaderboards struct {
	MostCredits         []MostCredit         `json:"mostCredits"`
	MostSubmittedCharts []MostSubmittedChart `json:"mostSubmittedCharts"`
}

type MostCredit struct {
	AgentSymbol string `json:"agentSymbol"`
	Credits     int64  `json:"credits"`
}

type MostSubmittedChart struct {
	AgentSymbol string `json:"agentSymbol"`
	ChartCount  int64  `json:"chartCount"`
}

type Link struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ServerResets struct {
	Next      string `json:"next"`
	Frequency string `json:"frequency"`
}

type Stats struct {
	Agents    int64 `json:"agents"`
	Ships     int64 `json:"ships"`
	Systems   int64 `json:"systems"`
	Waypoints int64 `json:"waypoints"`
}

type NewAgent struct {
	Data NewAgentData `json:"data"`
}

type NewAgentData struct {
	Agent    Agent    `json:"agent"`
	Contract Contract `json:"contract"`
	Faction  Faction  `json:"faction"`
	Ship     Ship     `json:"ship"`
	Token    string   `json:"token"`
}

type Agent struct {
	AccountID       string `json:"accountId"`
	Symbol          string `json:"symbol"`
	Headquarters    string `json:"headquarters"`
	Credits         int64  `json:"credits"`
	StartingFaction string `json:"startingFaction"`
	ShipCount       int64  `json:"shipCount"`
}

type Contract struct {
	ID               string    `json:"id"`
	FactionSymbol    string    `json:"factionSymbol"`
	Type             string    `json:"type"`
	Terms            Terms     `json:"terms"`
	Accepted         bool      `json:"accepted"`
	Fulfilled        bool      `json:"fulfilled"`
	Expiration       time.Time `json:"expiration"`
	DeadlineToAccept time.Time `json:"deadlineToAccept"`
}

type Terms struct {
	Deadline time.Time `json:"deadline"`
	Payment  Payment   `json:"payment"`
	Deliver  []Deliver `json:"deliver"`
}

type Deliver struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int64  `json:"unitsRequired"`
	UnitsFulfilled    int64  `json:"unitsFulfilled"`
}

type Payment struct {
	OnAccepted  int64 `json:"onAccepted"`
	OnFulfilled int64 `json:"onFulfilled"`
}

type Faction struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Headquarters string  `json:"headquarters"`
	Traits       []Trait `json:"traits"`
	IsRecruiting bool    `json:"isRecruiting"`
}

type Trait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Units       *int64 `json:"units,omitempty"`
}

type Ship struct {
	Symbol       string       `json:"symbol"`
	Registration Registration `json:"registration"`
	Nav          Nav          `json:"nav"`
	Crew         Crew         `json:"crew"`
	Frame        Frame        `json:"frame"`
	Reactor      Engine       `json:"reactor"`
	Engine       Engine       `json:"engine"`
	Cooldown     Cooldown     `json:"cooldown"`
	Modules      []Module     `json:"modules"`
	Mounts       []Mount      `json:"mounts"`
	Cargo        Cargo        `json:"cargo"`
	Fuel         Fuel         `json:"fuel"`
}

type Cargo struct {
	Capacity  int64   `json:"capacity"`
	Units     int64   `json:"units"`
	Inventory []Trait `json:"inventory"`
}

type Cooldown struct {
	ShipSymbol       string    `json:"shipSymbol"`
	TotalSeconds     int64     `json:"totalSeconds"`
	RemainingSeconds int64     `json:"remainingSeconds"`
	Expiration       time.Time `json:"expiration"`
}

type Crew struct {
	Current  int64  `json:"current"`
	Required int64  `json:"required"`
	Capacity int64  `json:"capacity"`
	Rotation string `json:"rotation"`
	Morale   int64  `json:"morale"`
	Wages    int64  `json:"wages"`
}

type Engine struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Condition    int64        `json:"condition"`
	Integrity    int64        `json:"integrity"`
	Speed        *int64       `json:"speed,omitempty"`
	Requirements Requirements `json:"requirements"`
	PowerOutput  *int64       `json:"powerOutput,omitempty"`
}

type Requirements struct {
	Power int64 `json:"power"`
	Crew  int64 `json:"crew"`
	Slots int64 `json:"slots"`
}

type Frame struct {
	Symbol         string       `json:"symbol"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Condition      int64        `json:"condition"`
	Integrity      int64        `json:"integrity"`
	ModuleSlots    int64        `json:"moduleSlots"`
	MountingPoints int64        `json:"mountingPoints"`
	FuelCapacity   int64        `json:"fuelCapacity"`
	Requirements   Requirements `json:"requirements"`
}

type Fuel struct {
	Current  int64    `json:"current"`
	Capacity int64    `json:"capacity"`
	Consumed Consumed `json:"consumed"`
}

type Consumed struct {
	Amount    int64     `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type Module struct {
	Symbol       string       `json:"symbol"`
	Capacity     int64        `json:"capacity"`
	Range        int64        `json:"range"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Requirements Requirements `json:"requirements"`
}

type Mount struct {
	Symbol       string       `json:"symbol"`
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	Strength     int64        `json:"strength"`
	Deposits     []string     `json:"deposits"`
	Requirements Requirements `json:"requirements"`
}

type Nav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          Route  `json:"route"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type Route struct {
	Destination   Destination `json:"destination"`
	Origin        Destination `json:"origin"`
	DepartureTime time.Time   `json:"departureTime"`
	Arrival       time.Time   `json:"arrival"`
}

type Destination struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int64  `json:"x"`
	Y            int64  `json:"y"`
}

type Registration struct {
	Name          string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role          string `json:"role"`
}

type ListFactions struct {
	Data []Datum `json:"data"`
	Meta Meta    `json:"meta"`
}

type Datum struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Headquarters string  `json:"headquarters"`
	Traits       []Trait `json:"traits"`
	IsRecruiting bool    `json:"isRecruiting"`
}

type Meta struct {
	Total int64 `json:"total"`
	Page  int64 `json:"page"`
	Limit int64 `json:"limit"`
}

type GetFaction struct {
	Data GetFactionData `json:"data"`
}

type GetFactionData struct {
	Symbol       string  `json:"symbol"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Headquarters string  `json:"headquarters"`
	Traits       []Trait `json:"traits"`
	IsRecruiting bool    `json:"isRecruiting"`
}
