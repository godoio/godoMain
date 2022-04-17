<script setup>
import { RouterLink, RouterView } from 'vue-router'
import HelloWorld from '@/components/HelloWorld.vue'
</script>
<template>
<div>
  <p class="hide" id="modalOp">
    <Modal 
      :taskIdProp="taskIdProp" 
      :taskTitleProp="taskTitleProp" 
      :taskDescriptionProp="taskDescriptionProp" 
      :tasksProp="tasksProp" 
      @edited="editTask" 
      @deleted="deleteTask" />
  </p>
  <p>Tasks are published here</p>
  
  <div class="container">
    <ul>
      <li v-for="task in tasks" :key="task.TaskId">
        <button v-on:click="getModal(task.TaskId, task.TaskTitle, task.TaskDescription)" class="card">
          <h2>{{task.TaskId}} {{task.TaskTitle}}</h2>
          <p id="desc">{{task.TaskDescription}}</p>
          <p id="date">{{task.TaskDate}}</p>
        </button>
      </li>
    </ul>
  </div>

  <ul >
    <li>
      <div class="form-card search-card">
        <ul>
          <li class="container-spacing">
            <span class="floated"><input type="text" placeholder="Enter task ID" v-model="taskId"></span>
          </li>
          <li class="container-spacing">
            <span>
              <button v-on:click="getTask">Search by Id</button>
              <button v-on:click="reload" style="margin-left:5px;">Refresh</button>
            </span>
          </li>
        </ul>
      </div>
    </li>

    <li>
      <div class="form-card create-card">
        <ul>
          <li class="container-spacing">
            <span class="floated"><input placeholder="Task title" class="input-size" type="text" v-model="taskTitle"></span><br>
            <span class="floated"><textarea placeholder="Task description" type="text" class="text-size" v-model="taskDesc"></textarea></span>
          </li>
          <li class="container-spacing">
            <span>
              <button v-on:click="createTask">Create Task</button>
            </span>
          </li>
        </ul>
      </div>
    </li>
  </ul>  
</div>
  
</template>

<script>
import axios from 'axios';
import Modal from './components/Modal.vue';

export default {
  setup() {
  },
  components() {
    Modal
  },
  data() {
    return {
      tasks : null,
      taskId : null,
      taskTitle : null,
      taskDesc: null,
      taskIdProp: null,
      taskTitleProp: null,
      taskDescriptionProp: null,
      tasksProp : null
    };
  },
  methods: {
  getTask: async function() {
    const response = await axios.get('http://localhost:4000/view?id='+this.taskId);
    this.tasks = [];
    this.tasks.push(response.data);
  },
  reload: async function() {
    const response = await axios.get('http://localhost:4000/viewAll');
    this.taskId = null
    this.tasks = response.data;
  },
  createTask: async function() {
    const response = await axios.get('http://localhost:4000/create?title='+this.taskTitle+'&desc='+this.taskDesc);
    this.taskTitle = '';
    this.taskDesc = '';
    this.tasks.push(response.data);
  },
  getModal: function(taskId, taskTitle, taskDescription) {
    var element = document.getElementById("modalOp");
      this.taskIdProp = taskId;
      this.taskTitleProp = taskTitle;
      this.taskDescriptionProp = taskDescription;
      this.tasksProp = this.tasks;
      element.style.visibility = "visible";
  },
  editTask: function(data) {
    this.tasks.find(task => task.TaskId === data.TaskId).TaskTitle = data.TaskTitle;
    this.tasks.find(task => task.TaskId === data.TaskId).TaskDescription = data.TaskDescription;
    this.tasks.find(task => task.TaskId === data.TaskId).TaskDate = data.TaskDate;
  },
  deleteTask: function(data) {
    this.tasks = data;
  }
  },
  async created() {
    const response = await axios.get('http://localhost:4000/viewAll');
    this.tasks = response.data;
    var element = document.getElementById("modalOp");
    element.style.visibility = "hidden";
  }, 

};
</script>

<style>
ul {
  list-style-type: none;
}
.card {
  background-color: white;
  border: none;
  box-shadow: 0 10px 20px 0 rgba(0.2,0,0.2,0.5);
  width: 930px;
  border-radius: 15px; /* 5px rounded corners */
  text-align: center;
  font-family: Georgia;
  margin-left: 30px;
  margin-bottom: 1rem;
  box-decoration-break: slice;
  padding-top: 10px;
  padding-bottom: 10px;
  padding-left: 10px;
  padding-right: 10px;
  position: relative;
  cursor: pointer;
}
.container {
  display: flex inline;
  flex-flow: row;
  justify-content: space-between;
}
#date {
  font-size: 75%;
  text-align: right;
  padding-right: 40px;
}
.form-card {
  box-shadow: 0 10px 20px 0 rgba(0.2,0,0.2,0.5);
  transition: 0.3s;
  width: 300px;
  height: 200px;
  text-align: center;
  display: flex inline;
  font-family: Georgia;
  box-decoration-break: slice;
  justify-content: space-between;
  padding-top: 10px;
  padding-right: 50px;
  padding-left: 10px;
  border-radius: 15px; /* 5px rounded corners */
  margin-left: 30px;
  margin-bottom: 1rem;
}
.floated {
  float:left;
}
.container-spacing {
  padding-top: 10px;
  margin-bottom: 1rem;
}
.hide {
  visibility: hidden;
}
.search-card {
  height:100px;
  width: 200px;
  top:200px;
  right: 200px;
  position: absolute;
}
.create-card {
  height:200px;
  width: 300px;
  top:320px;
  right: 100px;
  position: absolute;
}
.text-size {
  height:100px;
  width:250px;
  font-family: Georgia;
}
.input-size {
  font-family: Georgia;
  text-align: center;
}
.desc {
  font-size: 110%;
  font-family: Georgia;
}
</style>
