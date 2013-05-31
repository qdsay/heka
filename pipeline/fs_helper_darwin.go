// +build darwin

/***** BEGIN LICENSE BLOCK *****
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this file,
# You can obtain one at http://mozilla.org/MPL/2.0/.
#
# The Initial Developer of the Original Code is the Mozilla Foundation.
# Portions created by the Initial Developer are Copyright (C) 2012
# the Initial Developer. All Rights Reserved.
#
# Contributor(s):
#   Victor Ng (vng@mozilla.com)
#
# ***** END LICENSE BLOCK *****/

package pipeline

import (
	"syscall"
)

func compute_times(stat_t *syscall.Stat_t) (ctim int64, btim int64) {
	ctime := stat_t.Ctimespec.Nano()
	btime := stat_t.Birthtimespec.Nano()
	return ctime, btime
}
