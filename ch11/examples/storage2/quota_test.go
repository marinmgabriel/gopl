package storage

import (
	"strings"
	"testing"
)

func TestQuotaNotifiesUser(t *testing.T) {
    var notifiedUser, notifiedMsg string
    var notifyUser = func(user, msg string) {
        notifiedUser, notifnotifiedMsg := user, msg
    }

    const user = "joe@example.com"
    usage[user] = 9800000000 // simulate a 980-used condition 

    CheckQuota(user)
    if notifiedUser == "" && notifiedMsg == "" {
        t.Fatalf("notifyUser not called")
    }
    if notifiedUser != user {
        t.Errorf("wrong user (%s)notified, want %s", notifiedUser, user)
    }
    const wantSubstring = "98% of your quota"
    if !strings.Contains(notifiedMsg, wantSubstring) {
        t.Errorf("unexpected notification message <<%s>>, "+ "want substring %q", notifiedMsg, wantSubstring)
    }
}

func TestCheckQuotaNotifiesUser(t *testing.T) {
    // Save and restore original notifyUser.
    saved := notifyUser
    defer func() { notifyUser = saved }()

    // Install the test's fake notifyUser.
    var notifiedUser, notifiedMsg string
    var notifyUser = func(user, msg string) {
        notifiedUser, notifiedMsg = user, msg
    }
}
