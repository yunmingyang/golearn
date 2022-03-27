package storage

import (
	"strings"
	"testing"
)



func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	// In testing, just record notifiedUser and notifiedMsg, and saved the original function to not
	// effect the follow test
	saved := notifyUser
	defer func () { notifyUser = saved }()
	notifyUser = func (username, msg string)  {
		notifiedUser, notifiedMsg = username, msg
	}
	// ...simulate a 980MB-used condition
	bytesInUseMock = 980000000

	const user = "joe@example.org"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifiedUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s", notifiedUser, user)
	}

	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf(
			"unexpected notification message <<%s>>, want substring %s",
			notifiedMsg,
			wantSubstring,
		)
	}
}