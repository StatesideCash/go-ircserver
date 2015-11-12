package main

import ()

/// TODO Double and triple check that ALL codes are right!!!
/// TODO Message commands(Sections 4 and 5)
/// TODO Command responses(Sections 6.2 and 6.3)

/// Maintains the state of the user from given commands
type User struct {
	Username        string
	Realname        string
	IsRegistered    bool
	IsOper          bool
	Nick            string
	SessionPassword string
	QuitMessage     string
}

/// Keeps current state of the server
type Server struct {
	Users        []*User
	AllowedOpers map[string]string
}

/// For holding string tokens from messages sent from the client
type Message struct {
	Prefix  string
	Command string
	Params  string
}

/// Indicates success
const SUCCESS = 0

/// ERROR MESSAGES
const (
	ERR_NOSUCHNICK        = 401
	ERR_NOSUCHSERVER      = 402
	ERR_NOSUCHCHANNEL     = 403
	ERR_CANNOTSENDTOCHAN  = 404
	ERR_TOOMANYCHANNELS   = 405
	ERR_WASNOSUCHNICK     = 406
	ERR_TOOMANYTARGET     = 407
	ERR_NOORIGIN          = 409
	ERR_NORECIPIENT       = 411
	ERR_NOTTEXTTOSEND     = 412
	ERR_NOTTOPLEVEL       = 413
	ERR_WILDTOPLEVEL      = 414
	ERR_UNKNOWNCOMMAND    = 421
	ERR_NOMOTD            = 422
	ERR_NOADMININFO       = 423
	ERR_FILEERROR         = 424
	ERR_NONICKNAMEGIVE    = 431
	ERR_ERONEOUSNICKNAME  = 432
	ERR_NICKNAMEINUSE     = 433
	ERR_NICKCOLLISION     = 436
	ERR_USERNOTINCHANNEL  = 441
	ERR_NOTONCHANNEL      = 442
	ERR_USERONCHANNEL     = 443
	ERR_NOLOGIN           = 444
	ERR_SUMMONDISABLED    = 445
	ERR_USERSDISABLED     = 446
	ERR_NOTREGISTERED     = 451
	ERR_NEEDMOREPARAMS    = 461
	ERR_ALREADYREGISTERED = 462
	ERR_NOPERMFORHOST     = 463
	ERR_PASSWDMISMATCH    = 464
	ERR_YOUREBANNEDCREEP  = 465
	ERR_KEYSET            = 467
	ERR_CHANNELISFULL     = 471
	ERR_UNKNOWNMODE       = 472
	ERR_INVITEONLYCHAN    = 473
	ERR_BANNEDFROMCHAN    = 474
	ERR_BADCHANNELKEY     = 475
	ERR_NOPRIVILEGES      = 481
	ERR_CHANOPRIVSNEEDED  = 482
	ERR_CANTKILLSERVER    = 483
	ERR_NOOPERHOST        = 491
	ERR_UMODEUNKNOWNFLAG  = 501
	ERR_USERSDONTMATCH    = 502
)

const (
	RPL_YOUREOPER = 1 ///TODO Give value
	EXIT_STATUS   = -1
)
