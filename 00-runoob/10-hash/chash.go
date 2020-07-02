package main

import . "fmt"
import . "strconv"

func main(){
  hashs()
}

func hashs(){
  getInstance()
  
  put("a","a_out");Print(get("a"));dsp()
  put("b","b_out");Print(get("b"));dsp()
  put("p","p_out");Print(get("p"));dsp()
}

type HashMap struct {
  key string
  value string
  hashCode int
  next *HashMap
}

var table [16](*HashMap)

func initTable(){ for i:= range table{ table[i] = &HashMap{"","",i,nil} } }

func dsp() {
  var count int = 0;
  for i := range table {
    if len(table[i].key) > 0{
      Print("{" + table[i].key + "," + table[i].value + "," + Itoa(table[i].hashCode) + "}")
	  count++
	}
  }
  Println(" (",count,")");
}

func getInstance() [16](*HashMap){ 
  if(table[0] == nil){ initTable() }
  return table
}

func genHashCode(k string) int{
  if(len(k) == 0) { return 0 }

  var hashCode int = 0
  var lastIndex int = len(k) - 1
  for i := range k {
    if i == lastIndex { hashCode += int(k[i]); break }
	hashCode += (hashCode + int(k[i])) * 31
  }

  return hashCode
}

func indexTable(hashCode int) int { return hashCode % 6 }
func indexNode(hashCode int) int { return hashCode >> 4 }

func put(k string, v string) string{
  var hashCode = genHashCode(k)
  var thisNode = HashMap{k,v,hashCode,nil}
  
  var tableIndex = indexTable(hashCode)
  var nodeIndex = indexNode(hashCode)
  
  var headPtr [16](*HashMap) = getInstance()
  var headNode = headPtr[tableIndex]
  
  if(*headNode).key == "" { *headNode = thisNode; return "" }
  
  var lastNode *HashMap =  headNode
  var nextNode *HashMap = (*headNode).next
  
  for nextNode != nil && (indexNode((*nextNode).hashCode) < nodeIndex){
    lastNode = nextNode; nextNode = (*nextNode).next
  }
  if (*lastNode).hashCode == thisNode.hashCode{
    var oldValue string = lastNode.value
	lastNode.value = thisNode.value
	return oldValue
  }
  if lastNode.hashCode < thisNode.hashCode { lastNode.next = &thisNode }
  if nextNode != nil { thisNode.next = nextNode }
  
  return ""
}

func get(k string) string{
  var hashCode = genHashCode(k)
  var tableIndex = indexTable(hashCode)
  
  var headPtr [16](*HashMap) = getInstance()
  var node *HashMap = headPtr[tableIndex]
  
  if (*node).key == k { return (*node).value }
  
  for (*node).next != nil { if k == (*node).key { return (*node).value } }
	
  return ""
}
