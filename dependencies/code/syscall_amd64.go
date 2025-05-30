//go:build linux && amd64

package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
)

const (
	SYS_GETRANDOM = 318
	SYS_RSEQ      = 334
	SYS_SENDMMSG  = 307
)

var ALLOW_SYSCALLS = []int{
	// file io
	syscall.SYS_NEWFSTATAT, syscall.SYS_IOCTL, syscall.SYS_LSEEK, syscall.SYS_GETDENTS64,
	syscall.SYS_WRITE, syscall.SYS_CLOSE, syscall.SYS_OPENAT, syscall.SYS_READ,
	// thread
	syscall.SYS_FUTEX,
	// memory
	syscall.SYS_MMAP, syscall.SYS_BRK, syscall.SYS_MPROTECT, syscall.SYS_MUNMAP, syscall.SYS_RT_SIGRETURN,
	syscall.SYS_MREMAP,

	// user/group
	syscall.SYS_SETUID, syscall.SYS_SETGID, syscall.SYS_GETUID,
	// process
	syscall.SYS_GETPID, syscall.SYS_GETPPID, syscall.SYS_GETTID,
	syscall.SYS_EXIT, syscall.SYS_EXIT_GROUP,
	syscall.SYS_TGKILL, syscall.SYS_RT_SIGACTION, syscall.SYS_IOCTL,
	syscall.SYS_SCHED_YIELD,
	syscall.SYS_SET_ROBUST_LIST, syscall.SYS_GET_ROBUST_LIST, SYS_RSEQ,

	// time
	syscall.SYS_CLOCK_GETTIME, syscall.SYS_GETTIMEOFDAY, syscall.SYS_NANOSLEEP,
	syscall.SYS_EPOLL_CREATE1,
	syscall.SYS_EPOLL_CTL, syscall.SYS_CLOCK_NANOSLEEP, syscall.SYS_PSELECT6,
	syscall.SYS_TIME,

	syscall.SYS_RT_SIGPROCMASK, syscall.SYS_SIGALTSTACK, SYS_GETRANDOM,
}

var ALLOW_ERROR_SYSCALLS = []int{
	syscall.SYS_CLONE,
	syscall.SYS_MKDIRAT,
	syscall.SYS_MKDIR,
}

var ALLOW_NETWORK_SYSCALLS = []int{
	syscall.SYS_SOCKET, syscall.SYS_CONNECT, syscall.SYS_BIND, syscall.SYS_LISTEN, syscall.SYS_ACCEPT, syscall.SYS_SENDTO, syscall.SYS_RECVFROM,
	syscall.SYS_GETSOCKNAME, syscall.SYS_RECVMSG, syscall.SYS_GETPEERNAME, syscall.SYS_SETSOCKOPT, syscall.SYS_PPOLL, syscall.SYS_UNAME,
	syscall.SYS_SENDMSG, SYS_SENDMMSG, syscall.SYS_GETSOCKOPT,
	syscall.SYS_FSTAT, syscall.SYS_FCNTL, syscall.SYS_FSTATFS, syscall.SYS_POLL, syscall.SYS_EPOLL_PWAIT,
}

func main() {
	var syscalls []int
	syscalls = append(syscalls, SYS_GETRANDOM, SYS_RSEQ, SYS_SENDMMSG)
	syscalls = append(syscalls, ALLOW_SYSCALLS...)
	syscalls = append(syscalls, ALLOW_ERROR_SYSCALLS...)
	syscalls = append(syscalls, ALLOW_NETWORK_SYSCALLS...)
	syscallsStr := make([]string, len(syscalls))
	for i, syscall := range syscalls {
		syscallsStr[i] = strconv.Itoa(syscall)
	}
	result := strings.Join(syscallsStr, ",")
	fmt.Println(result)

}
