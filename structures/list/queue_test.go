package list

import "testing"

func GenericQueueEnqueueDequeue(q Queue) string {
    q.Enqueue(5)
    q.Enqueue(6)
    q.Enqueue(7)

    item := q.Dequeue()
    if item != 5 {
        return "Dequeuing items from the queue has failed."
    }

    item = q.Dequeue()
    if item != 6 {
        return "Dequeued item was not properly removed from the queue."
    }

    return ""
}

func TestSinglyQueueEnqueueDequeue(t *testing.T) {
    s := &Singly{}
    e := GenericQueueEnqueueDequeue(s)
    if len(e) > 0 {
        t.Error(e)
    }
}

func TestDoublyQueueEnqueueDequeue(t *testing.T) {
    d := &Doubly{}
    e := GenericQueueEnqueueDequeue(d)
    if len(e) > 0 {
        t.Error(e)
    }
}
