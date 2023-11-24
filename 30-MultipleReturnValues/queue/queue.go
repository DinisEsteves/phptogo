package queue

type job struct {
	body    string
	timeout int
}

var queue []job

func Enqueue(body string, timeout int) (jobTimeout int, j job) {
	j = job{
		body:    body,
		timeout: timeout,
	}

	queue := append(queue, j)

	jobTimeout = len(queue)

	return
}
