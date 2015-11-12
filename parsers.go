package main

import (
	"bufio"
	"strings"
)

/// Please note that all functions prefixed as Handle are mutators if required

/// This scans the sent message and turns it into a struct with the
/// basic parts of the message sepereated into a Message struct. This only
/// builds the struct, handling should be done on the returned struct
/// with the given handling functions
func ParseMessage(message string) Message {
	/// Initialize the scanner for parsing the text split by whitespace
	scanner := bufio.NewScanner(strings.NewReader(message))
	scanner.Split(bufio.ScanWords)

	/// Instantiate the Message to store the data in
	msg := Message{}

	wasPrefix := true

	/// Prefix
	if !scanner.Scan() {
		return msg
	} else {
		/// Read in the first word prepended by a ":" as the prefix
		/// The prefix is used by servers to identify the true source of a message
		text := scanner.Text()
		if string(text[0]) == ":" {
			msg.Prefix = text
			/// TODO Add in the optional USER and HOST fields
		} else {
			wasPrefix = false
		}
	}

	/// Command
	if wasPrefix {
		if !scanner.Scan() {
			return msg
		} else {
			/// The next word is the command
			command := scanner.Text()
			msg.Command = command
		}
	} else {
		msg.Command = scanner.Text()
	}

	/// Params
	scanner.Split(bufio.ScanLines)
	if !scanner.Scan() {
		return msg
	} else {
		/// The rest of the line is the params
		params := scanner.Text()
		msg.Params = params
	}
	return msg
}

/// Decides which function to call to handle a given command
/// Returns an int for the status to send to the user
func HandleCommand(msg Message, user *User, server *Server) int {
	if msg.Command == "PASS" {
		return HandlePassword(msg.Params, user)
	} else if msg.Command == "NICK" {
		return HandleNick(msg.Params, user, server)
	} else if msg.Command == "USER" {
		return HandleUser(msg.Params, user, server)
	} else if msg.Command == "OPER" {
		return HandleOper(msg.Params, user, server)
	} else if msg.Command == "QUIT" {
		return HandleQuit(msg.Params, user, server)
	}

	/// If we get here, the command was invalid
	return ERR_UNKNOWNCOMMAND
}

/// Sets the user's password for the session.
/// Returns SUCCESS, ERR_ALREADYREGISTRED or ERR_NEEDMOREPARAMS
func HandlePassword(params string, user *User) int {
	if len(params) == 0 {
		return ERR_NEEDMOREPARAMS
	} else if user.IsRegistered {
		return ERR_ALREADYREGISTERED
	}
	user.SessionPassword = params
	return SUCCESS
}

/// Sets or changes the user's password
/// Returns SUCCESS, ERR_ERRONEOUSNICKNAME, ERR_NONICKNAMEGIVEN, ERR_NICKCOLLISION, AND ERR_NICKNAMEINUSE
func HandleNick(params string, user *User, server *Server) int {
	if len(params) == 0 {
		return ERR_NONICKNAMEGIVE
	}
	/// TODO Check for ERR_ERRONEOUSNICKNAME and ERR_NICKCOLLISION

	/// Only the first word is to be the nickname
	nick := strings.Split(params, " ")[0]

	for _, u := range server.Users {
		if u.Nick == nick {
			return ERR_NICKNAMEINUSE
		}
	}
	user.Nick = nick
	return SUCCESS
}

/// Handles the USER command. Currently ignores the hostname and servername
/// args because this software does not support multiple servers
func HandleUser(params string, user *User, server *Server) int {
	if user.IsRegistered {
		return ERR_ALREADYREGISTERED
	}

	var username, realname string
	scanner := bufio.NewScanner(strings.NewReader(params))
	scanner.Split(bufio.ScanWords)

	count := 0
	if scanner.Scan() {
		username = scanner.Text()
		count++
		if scanner.Scan() {
			//TODO hostname = scanner.Text()
			count++
			if scanner.Scan() {
				//TODO servername = scanner.Text()
				count++
				scanner.Split(bufio.ScanLines)
				if scanner.Scan() {
					realname = scanner.Text()
					count++
				}
			}
		}
	}
	if count != 4 {
		return ERR_NEEDMOREPARAMS
	}

	user.Username = username
	user.Realname = realname
	return SUCCESS
}

/// Tries to give the user OP status
func HandleOper(params string, user *User, server *Server) int {
	args := strings.Split(params, " ")
	if len(args) < 2 {
		return ERR_NEEDMOREPARAMS
	}
	username := args[0]
	password := args[1]

	if user.IsOper {
		return RPL_YOUREOPER
	}

	if pass, ok := server.AllowedOpers[username]; ok {
		if pass == password {
			user.IsOper = true
			return SUCCESS
		} else {
			return ERR_PASSWDMISMATCH
		}
	}
	return ERR_NOOPERHOST
}

func HandleQuit(params string, user *User, server *Server) int {
	if len(params) > 0 {
		user.QuitMessage = params
	}
	return EXIT_STATUS
}
