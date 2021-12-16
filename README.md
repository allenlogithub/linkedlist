# linkedlist
Declaration:
```
// Singly Linked List
l := linkedlist.LinkedList{}
// Doubly Linked List
l = linkedlist.DoublyLinkedList{} 
```
Push(el):  
&nbsp; Add an element at the end of the linkedlist
```
l.Push(3)
```
PushAt(pos, el):  
&nbsp; Add an element at a given position of the linkedlist
```
l.PushAt(1, 1)
```
Pop(el):  
&nbsp; Remove the first element:el
```
l.Pop(1)
```
PopAt(pos):  
&nbsp; Remove an element:el at the given position of the linkedlist
```
l.PopAt(0)
```
InsertBefore(*node, el)/ InsertAfter(*node, el):  
&nbsp; Insert an element before/ after a node, support DLL at the moment  
```
l.InsertAfter(l.Head.Next, 1)
l.InsertBefore(l.Head.Next, 1)
```
Show():  
&nbsp; Show the linkedlist info
```
l.Show()
// if Doubly Linked List, l.Show(bool), true for traversing from Head
```

Error message:
```
if err := l.Pop(1); err != nil {
  fmt.Println(err.Error)
  fmt.Println(err.Message)
}
```
