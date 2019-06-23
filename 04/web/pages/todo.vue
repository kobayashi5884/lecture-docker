<template>
  <div>
    <h2>TODO List</h2>

    <div class="new-todo">
      <!-- 
        v-modelは、dataから値を参照し、inputの値が変更されると、dataに反映します。
        要するに、inputの値とdataの値が同期します。 
      -->
      <input 
        v-model="newTodo" 
        type="text"
      >
      <!-- v-onは、methodsで定義された関数を呼び出します。 -->
      <button v-on:click="add">add</button>
    </div>

    <ul>
      <!-- 
        v-forは、dataの配列（ここではtodos）の要素に応じて、
        DOMの要素を描画することができます。 
      -->
      <!-- v-bindは、（文字列ではなく）変数を参照する場合に使用します。 -->
      <li 
        v-for="(todo, i) in todos" 
        v-bind:key="i"
        class="todos"
      >
        <!-- v-if、v-else-if、v-elseは、DOM要素の表示、非表示を制御します。 -->
        <!-- 
          v-showというディレクティブもあります。
          頻繁に表示、非表示を切り替えるような場合には、v-showの方が効率がよい場合があります。 
        -->
        <span 
          v-if="!todo.editable"
          key="default"
        >
          {{ todo.text }}
          <button v-on:click="edit(i)">
            edit
          </button>
          <button v-on:click="todos.splice(i, 1)">
            delete
          </button>
        </span>
        
        <span 
          v-else
          key="edit"
        >
          <input 
            v-model="todo.newText"
            type="text"
          >
          <button v-on:click="cancel(i)">cancel</button>
          <button v-on:click="update(i)">update</button>
        </span>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newTodo : '',
      todos: [
        { text: 'Shopping', editable: false, newText: '' },
        { text: 'Cooking',  editable: false, newText: '' },
        { text: 'Homework', editable: false, newText: '' },
      ],
    }
  },
  methods: {
    add() {
      if (!this.newTodo) {
        alert('Text is empty')
        return
      }
      this.todos.push({ 
        text: this.newTodo,
        editable: false,
        newText: '', 
      })
    },
    edit(i) {
      // オブジェクトの配列を更新する場合には、
      // オブジェクトのプロパティを更新するのではなく、
      // オブジェクトを丸ごと入れ替える必要があります。
      const newTodo = JSON.parse(JSON.stringify(this.todos[i]))
      newTodo.newText = newTodo.text
      newTodo.editable = true
      this.todos.splice(i, 1, newTodo)
      // this.$set(this.todos, i, newTodo) というVue.js固有の書き方もあります。
    },
    cancel(i) {
      const newTodo = JSON.parse(JSON.stringify(this.todos[i]))
      newTodo.editable = false
      this.todos.splice(i, 1, newTodo)
    },
    update(i) {
      const newTodo = JSON.parse(JSON.stringify(this.todos[i]))
      newTodo.text = newTodo.newText
      newTodo.editable = false
      this.todos.splice(i, 1, newTodo)
    },
  },
}
</script>

<style>
.new-todo {
  margin-bottom: 16px;
}
.todos {
  margin-bottom: 8px;
}
</style>
