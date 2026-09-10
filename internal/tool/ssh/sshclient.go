package ssh

import (
    "context"
	"fmt"
	"net"
	"strings"
	"time"

	"golang.org/x/crypto/ssh"
)


type SSHClient struct {
    sshclient *ssh.Client
}



func RunSSHCommand(ctx context.Context, host, user, password, command string, port ...string) (string, error) {
    addr := host
	if !strings.Contains(host, ":") {
		if port != nil {
            addr = host + port[0]
        } else {
            addr = host + ":22"
        }
	}
    
    client, err := dial(ctx, addr, user, password)
	if err != nil {
		return "", fmt.Errorf("ssh %s: %w", addr, err)
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("new session: %w", err)
	}
	defer session.Close()

    return combinedOutput(ctx, session, command)
}

func dial(ctx context.Context, addr, user, password string) (*ssh.Client, error) {
	cfg := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
			// Some servers are configured for keyboard-interactive
			// instead of plain password auth; answer with the same
			// password.
			ssh.KeyboardInteractive(func(_, _ string, questions []string, _ []bool) ([]string, error) {
				answers := make([]string, len(questions))
				for i := range answers {
					answers[i] = password
				}
				return answers, nil
			}),
		},
		// Demo environment: accept whatever host key the server presents
		// (equivalent to StrictHostKeyChecking=no).
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         10 * time.Second,
	}

	d := net.Dialer{Timeout: cfg.Timeout}
	conn, err := d.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("dial: %w", err)
	}

	c, chans, reqs, err := ssh.NewClientConn(conn, addr, cfg)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("handshake/auth: %w", err)
	}
	return ssh.NewClient(c, chans, reqs), nil
}


func combinedOutput(ctx context.Context, session *ssh.Session, cmd string) (string, error) {
	type result struct {
		out []byte
		err error
	}
	ch := make(chan result, 1)
	go func() {
		out, err := session.CombinedOutput(cmd)
		ch <- result{out, err}
	}()

	select {
	case r := <-ch:
		if r.err != nil {
			return "", r.err
		}
		return string(r.out), nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
