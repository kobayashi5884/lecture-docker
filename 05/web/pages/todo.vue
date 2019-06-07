<template>
  <div>
    <h2>TODO List</h2>

    <div class="new-todo">
      <input 
        v-model="newTodo" 
        type="text"
      >
      <button v-on:click="add">add</button>
    </div>

    <div 
      v-for="(todo, i) in todos" 
      v-bind:key="todo.ID"
      class="todos"
    >
      <div 
        v-if="!todo.editable"
        key="default"
      >
        {{ todo.Text }}
        <button v-on:click="edit(i)">
          edit
        </button>
        <button v-on:click="remove(i)">
          delete
        </button>
      </div>
      
      <div 
        v-else
        key="edit"
      >
        <input 
          v-model="todo.newText"
          type="text"
        >
        <button v-on:click="cancel(i)">cancel</button>
        <button v-on:click="update(i)">update</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      newTodo : '',
      todos: [
        // APIから受け取るデータは、頭文字が大文字になっていますので、注意してください。
        // { ID: 1, Text: 'Shopping', editable: false, newText: '' },
      ],
    }
  },
  // asyncDataは、ページ遷移前にAPIからデータを取得し、dataに反映します。
  // 反映したいdataのオブジェクトを返り値にします。
  async asyncData({ $axios, error }) {
    let todos = []

    // APIにGETメソッドでリクエストを送り、todosを取得します。
    try {
      const { data } = await $axios.get('/todo')
      todos = data
    } catch (err) {
			error({
        statusCode: err.response.status,
        message:    err.response.statusText,
      })
    }

    // APIから取得するデータは、IDとTextだけですので、editableとnewTextを追加します。
    for (let i = 0; i < todos.length; i++) {
      todos[i] = { ...todos[i], editable: false, newText: '' }
    }

    // 取得したtodosをdataに反映します。
    return { todos }
  },
  methods: {
    async add() {
      if (!this.newTodo) {
        alert('Text is empty')
        return
      }

      let todo

      // APIにPOSTメソッドでリクエストを送り、todoを作成し、そのデータを取得します。
      try {
        const { data } = await this.$axios.post('/todo', { text: this.newTodo })
        todo = { ...data, editable: false, newText: '' }
      } catch (err) {
        alert('Failed to create a new item')  
        return
      }

      // 作成したtodoをtodosに追加します。
      this.todos.push(todo)
    },
    edit(i) {
      const newTodo = JSON.parse(JSON.stringify(this.todos[i]))
      newTodo.newText = newTodo.Text
      newTodo.editable = true
      this.todos.splice(i, 1, newTodo)
    },
    async remove(i) {
      // APIにDELETEメソッドでリクエストを送り、todoを削除します。
      try {
        await this.$axios.delete(`/todo/${this.todos[i].ID}`)
      } catch (err) {
        alert('Failed to delete the item')  
        return
      }

      this.todos.splice(i, 1)
    },
    cancel(i) {
      const newTodo = JSON.parse(JSON.stringify(this.todos[i]))
      newTodo.editable = false
      this.todos.splice(i, 1, newTodo)
    },
    async update(i) {
      let todo = JSON.parse(JSON.stringify(this.todos[i]))

      // APIにPUTメソッドでリクエストを送り、todoを更新し、そのデータを取得します。
      try {
        const { data } = await this.$axios.put(`/todo/${todo.ID}`, { text: todo.newText })
        todo = { ...data, editable: false, newText: '' }
      } catch (err) {
        alert('Failed to update the item')  
        return
      }

      // todoの更新結果をdataに反映します。
      this.todos.splice(i, 1, todo)
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
