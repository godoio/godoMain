<template>
    <div id="modalOp">
        <div className="modal">
            <div id="editMode" class="editMode">
                <input placeholder="Task title" class="input-size" type="text" v-model="taskTitleProp"><br>
                <textarea placeholder="Task description" type="text" class="text-size" v-model="taskDescriptionProp"></textarea>
            </div>
            <span>
              <button v-on:click="editTask()">{{editButton}}</button> 
              <button style="margin-left:5px;" v-on:click="deleteTask()">Delete Task</button>
              <button style="margin-left:5px;" v-on:click="closeModal()">Close Task</button>
            </span>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default {  
   name: 'Modal',
   data() {
       return {
           editButton : "Edit Task"
       }
   },
   props: { 
       taskIdProp: Number,
       taskTitleProp: String,
       taskDescriptionProp: String,
       tasksProp: Array
    },
    methods: {
        deleteTask : async function() {
            const response = await axios.get('http://localhost:4000/delete?id='+this.taskIdProp);
            var filtered = this.tasksProp.filter(function(value) {
                return value.TaskId != response.data.TaskId;
            });
            this.$emit('deleted', filtered);
            this.closeModal()
            
        },
        editTask : async function() {
            const response = await axios.get('http://localhost:4000/edit?title='+this.taskTitleProp+'&desc='+this.taskDescriptionProp+'&id='+this.taskIdProp);
            this.$emit('edited', response.data)
            this.closeModal()
        },
        closeModal : function() {
            var element = document.getElementById("modalOp");
            element.style.visibility = "hidden";
        }
    }
}  
</script>


<style>
.modal{
  border-radius: 10px;
  background-color: white;
  padding: 20px 20px 20px 20px;
  box-shadow: 0 10px 20px 0 rgba(0.2,0,0.2,0.5);
  justify-content: center;
  align-items: center;
  text-align: center;
  box-decoration-break: slice;
  z-index: 1000;
  position: absolute;
  width: 400px;
  height: 200px;
  display: flex inline-block;
  overflow-x: auto;
  top: 200px;
  left: 520px;
  justify-content: space-between;
  flex-flow: row;
  backdrop-filter: grayscale(0.5) opacity(0.8);
}
.editMode {
    display: flex inline-block;
}
.text-size {
  height:100px;
  width:250px;
}
.hide-modal {
  visibility: visible;
}
</style>