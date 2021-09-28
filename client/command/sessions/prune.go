package sessions

/*
	Sliver Implant Framework
	Copyright (C) 2021  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"context"

	"github.com/bishopfox/sliver/client/command/kill"
	"github.com/bishopfox/sliver/client/console"
	"github.com/bishopfox/sliver/protobuf/commonpb"
	"github.com/desertbit/grumble"
)

// SessionsPruneCmd - Forcefully kill stale sessions
func SessionsPruneCmd(ctx *grumble.Context, con *console.SliverConsoleClient) {
	sessions, err := con.Rpc.GetSessions(context.Background(), &commonpb.Empty{})
	if err != nil {
		con.PrintErrorf("%s\n", err)
		return
	}
	if len(sessions.GetSessions()) == 0 {
		con.PrintInfof("No sessions to prune\n")
		return
	}
	for _, session := range sessions.GetSessions() {
		if session.IsDead {
			con.Printf("Pruning session #%d ...", session.ID)
			err = kill.KillSession(session, true, con)
			if err != nil {
				con.Printf("failed!\n")
				con.PrintErrorf("%s\n", err)
			} else {
				con.Printf("done!\n")
			}
		}
	}
}
