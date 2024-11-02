package pointers

import (
	"testing"
)

func TestPushAndPopWithoutNil(t *testing.T) {
	queue := &Queue{}

	// Тестируем извлечение из пустой очереди
	value, ok := queue.PopWithoutNil()
	if ok {
		t.Errorf("Ожидалось, что PopWithoutNil вернет ok=false для пустой очереди, но получили ok=true")
	}
	if value != "" {
		t.Errorf("Ожидалось, что PopWithoutNil вернет пустую строку для пустой очереди, но получили %s", value)
	}

	// Добавляем элементы в очередь
	queue.Push("один")
	queue.Push("два")
	queue.Push("три")

	// Извлекаем элементы и проверяем порядок
	value, ok = queue.PopWithoutNil()
	if !ok {
		t.Errorf("Ожидалось, что PopWithoutNil вернет ok=true, но получили ok=false")
	}
	if value != "один" {
		t.Errorf("Ожидалось, что первый извлеченный элемент будет 'один', но получили %s", value)
	}

	value, ok = queue.PopWithoutNil()
	if !ok {
		t.Errorf("Ожидалось, что PopWithoutNil вернет ok=true, но получили ok=false")
	}
	if value != "два" {
		t.Errorf("Ожидалось, что второй извлеченный элемент будет 'два', но получили %s", value)
	}

	value, ok = queue.PopWithoutNil()
	if !ok {
		t.Errorf("Ожидалось, что PopWithoutNil вернет ok=true, но получили ok=false")
	}
	if value != "три" {
		t.Errorf("Ожидалось, что третий извлеченный элемент будет 'три', но получили %s", value)
	}

	// Теперь очередь должна быть пустой
	value, ok = queue.PopWithoutNil()
	if ok {
		t.Errorf("Ожидалось, что PopWithoutNil вернет ok=false для пустой очереди, но получили ok=true")
	}
	if value != "" {
		t.Errorf("Ожидалось, что PopWithoutNil вернет пустую строку для пустой очереди, но получили %s", value)
	}

	// Добавляем элемент после извлечения всех элементов
	queue.Push("четыре")
	value, ok = queue.PopWithoutNil()
	if !ok {
		t.Errorf("Ожидалось, что PopWithoutNil вернет ok=true, но получили ok=false")
	}
	if value != "четыре" {
		t.Errorf("Ожидалось, что извлеченный элемент будет 'четыре', но получили %s", value)
	}
}

func TestPushPopAlternating(t *testing.T) {
	queue := &Queue{}

	// Чередуем добавление и извлечение
	queue.Push("один")
	value, ok := queue.PopWithoutNil()
	if !ok || value != "один" {
		t.Errorf("Ожидалось, что будет извлечен 'один', но получили '%s'", value)
	}

	queue.Push("два")
	queue.Push("три")
	value, ok = queue.PopWithoutNil()
	if !ok || value != "два" {
		t.Errorf("Ожидалось, что будет извлечен 'два', но получили '%s'", value)
	}

	queue.Push("четыре")
	value, ok = queue.PopWithoutNil()
	if !ok || value != "три" {
		t.Errorf("Ожидалось, что будет извлечен 'три', но получили '%s'", value)
	}

	value, ok = queue.PopWithoutNil()
	if !ok || value != "четыре" {
		t.Errorf("Ожидалось, что будет извлечен 'четыре', но получили '%s'", value)
	}

	// Очередь должна быть пустой
	value, ok = queue.PopWithoutNil()
	if ok {
		t.Errorf("Ожидалось, что PopWithoutNil вернет ok=false для пустой очереди, но получили ok=true")
	}
}

func TestQueueEmptyAfterPops(t *testing.T) {
	queue := &Queue{}

	// Добавляем элементы
	queue.Push("один")
	queue.Push("два")

	// Извлекаем все элементы
	queue.PopWithoutNil()
	queue.PopWithoutNil()

	// Проверяем, что очередь пуста
	if queue.first != nil || queue.last != nil {
		t.Errorf("Ожидалось, что очередь будет пустой после извлечения всех элементов, но first или last не равны nil")
	}
}

func TestQueue_PopWithoutNil(t *testing.T) {
	type fields struct {
		first *QueueItem
		last  *QueueItem
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		want1  bool
	}{
		{
			name: "Pop from empty queue",
			fields: fields{
				first: nil,
				last:  nil,
			},
			want:  "",
			want1: false,
		},
		{
			name: "Pop from queue with one item",
			fields: fields{
				first: &QueueItem{
					value: "item1",
					next:  nil,
				},
				last: &QueueItem{
					value: "item1",
					next:  nil,
				},
			},
			want:  "item1",
			want1: true,
		},
		{
			name: "Pop from queue with multiple items",
			fields: fields{
				first: &QueueItem{
					value: "item1",
					next: &QueueItem{
						value: "item2",
						next:  nil,
					},
				},
				last: &QueueItem{
					value: "item2",
					next:  nil,
				},
			},
			want:  "item1",
			want1: true,
		},
		{
			name: "Pop until queue is empty",
			fields: fields{
				first: &QueueItem{
					value: "item1",
					next: &QueueItem{
						value: "item2",
						next: &QueueItem{
							value: "item3",
							next:  nil,
						},
					},
				},
				last: &QueueItem{
					value: "item3",
					next:  nil,
				},
			},
			want:  "item1",
			want1: true,
		},
		{
			name: "Pop with first element's next not nil",
			fields: fields{
				first: &QueueItem{
					value: "item1",
					next:  nil, // next is nil, but we won't set it to nil in PopWithoutNil
				},
				last: &QueueItem{
					value: "item1",
					next:  nil,
				},
			},
			want:  "item1",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &Queue{
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			got, got1 := queue.PopWithoutNil()
			if got != tt.want {
				t.Errorf("PopWithoutNil() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("PopWithoutNil() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueue_Push(t *testing.T) {
	type fields struct {
		first *QueueItem
		last  *QueueItem
	}
	type args struct {
		value string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantFirstValue string
		wantLastValue  string
		wantQueueSize  int
	}{
		{
			name: "Push into empty queue",
			fields: fields{
				first: nil,
				last:  nil,
			},
			args: args{
				value: "item1",
			},
			wantFirstValue: "item1",
			wantLastValue:  "item1",
			wantQueueSize:  1,
		},
		{
			name: "Push into non-empty queue",
			fields: fields{
				first: &QueueItem{
					value: "item1",
					next:  nil,
				},
				last: &QueueItem{
					value: "item1",
					next:  nil,
				},
			},
			args: args{
				value: "item2",
			},
			wantFirstValue: "item1",
			wantLastValue:  "item2",
			wantQueueSize:  1,
		},
		{
			name: "Push multiple items",
			fields: fields{
				first: &QueueItem{
					value: "item1",
					next: &QueueItem{
						value: "item2",
						next:  nil,
					},
				},
				last: &QueueItem{
					value: "item2",
					next:  nil,
				},
			},
			args: args{
				value: "item3",
			},
			wantFirstValue: "item1",
			wantLastValue:  "item3",
			wantQueueSize:  2,
		},
		{
			name: "Push empty string",
			fields: fields{
				first: nil,
				last:  nil,
			},
			args: args{
				value: "",
			},
			wantFirstValue: "",
			wantLastValue:  "",
			wantQueueSize:  1,
		},
		{
			name: "Push after queue was emptied",
			fields: fields{
				first: nil,
				last:  nil,
			},
			args: args{
				value: "new_item",
			},
			wantFirstValue: "new_item",
			wantLastValue:  "new_item",
			wantQueueSize:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			queue := &Queue{
				first: tt.fields.first,
				last:  tt.fields.last,
			}
			queue.Push(tt.args.value)

			// Проверяем значение первого элемента
			if queue.first == nil || queue.first.value != tt.wantFirstValue {
				t.Errorf("Push() first value = %v, want %v", queue.first.value, tt.wantFirstValue)
			}

			// Проверяем значение последнего элемента
			if queue.last == nil || queue.last.value != tt.wantLastValue {
				t.Errorf("Push() last value = %v, want %v", queue.last.value, tt.wantLastValue)
			}

			// Проверяем размер очереди
			size := 0
			for item := queue.first; item != nil; item = item.next {
				size++
			}
			if size != tt.wantQueueSize {
				t.Errorf("Push() queue size = %d, want %d", size, tt.wantQueueSize)
			}
		})
	}
}
