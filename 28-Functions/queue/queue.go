package queue

type job struct {
	body    string
	timeout int
}

var queue []job

const DefaultTimeout int = 5

func Enqueue(body string, timeout int) int {
	queue := append(queue, job{
		body:    body,
		timeout: timeout,
	})

	return len(queue)
}

func EnqueueWithoutTimeout(body string) int {
	return Enqueue(body, DefaultTimeout)
}
