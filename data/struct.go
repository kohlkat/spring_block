package main

type LedgerResponse struct {
	ID     int `json:"id"`
	Result struct {
		Ledger struct {
			Accepted            bool     `json:"accepted"`
			AccountHash         string   `json:"account_hash"`
			CloseFlags          int      `json:"close_flags"`
			CloseTime           int      `json:"close_time"`
			CloseTimeHuman      string   `json:"close_time_human"`
			CloseTimeResolution int      `json:"close_time_resolution"`
			Closed              bool     `json:"closed"`
			Hash                string   `json:"hash"`
			LedgerHash          string   `json:"ledger_hash"`
			LedgerIndex         string   `json:"ledger_index"`
			ParentCloseTime     int      `json:"parent_close_time"`
			ParentHash          string   `json:"parent_hash"`
			SeqNum              string   `json:"seqNum"`
			TotalCoins          string   `json:"totalCoins"`
			Total_Coins          string   `json:"total_coins"`
			TransactionHash     string   `json:"transaction_hash"`
			Transactions        []string `json:"transactions"`
		} `json:"ledger"`
		LedgerHash  string `json:"ledger_hash"`
		LedgerIndex int    `json:"ledger_index"`
		Validated   bool   `json:"validated"`
	} `json:"result"`
	Status string `json:"status"`
	Type   string `json:"type"`
}


type LedgerRequest struct {
	Id 				int `json:"id"`
	Command 		string `json:"command"`
	LedgerIndex 	string `json:"ledger_index"`
	Full 			bool `json:"full"`
	Accounts 		bool `json:"accounts"`
	Transactions 	bool `json:"transactions"`
	Expand 			bool `json:"expand"`
	OwnerFunds  	bool `json:"owner_funds"`
}



// Payload object
type message struct {
	Command string `json:"command"`
	//Taker     string    `json:"taker"`
	TakerGets takerGets `json:"taker_gets"`
	TakerPays takerPays `json:"taker_pays"`
	Limit     uint      `json:"limit"`
}

type takerGets struct {
	Currency string `json:"currency"`
}

type takerPays struct {
	Currency string `json:"currency"`
	Issuer   string `json:"issuer"`
}
