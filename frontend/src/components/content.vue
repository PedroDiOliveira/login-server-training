<template>
    <div class="content">
        <div class="input">
        <h1>Cadastro</h1>
        <form class="formulario" id="formulario" @submit.prevent="getFormValues">
            <div class="label">
                <label for="username">Usuario: </label>
                <input v-model="user.username" name="username" id="username" type="text" required>
            </div>
            <div class="label">
                <label for="idade">Idade: </label>
                <input v-model="user.idade" name="idade" id="idade" type="text" required>
            </div>    

            <div class="butao">
                <button class="butao" type="submit" v-on:click="getFormValues">Cadastrar</button>
            </div>
        </form>
    </div>
        <div class="dataBase">
            <h1>DataBase</h1>
            <div class="output" v-for="(cebola,indice) in lista"  >
                <h4><strong>Usuario</strong>{{indice + 1}}</h4>
                <p>Username: {{ cebola.username }}</p>
                <p>Idade: {{ cebola.idade }}</p>
            </div>
        </div>
    </div>    
</template>

<style scoped>

    .content{
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        color: white;
        padding: 3em;
        gap: 2em;
    }
    .input{
        display: flex;
        justify-content: center;
        align-items: center;
        flex-direction: column;
        width: 40%;
        height: 35em;
        border-radius: 10px;
        border: solid 2px #BDBFBF;
    }

    .formulario{
        display: flex;
        flex-direction: column;
        justify-content: space-between;
        padding: 4em;
        gap: 2em;
    }

    h1{
        font-size: 4ch;
        font-family: Impact, Haettenschweiler, 'Arial Narrow Bold', sans-serif;
    }

    label{
        font-size: xx-large;
        font-family:Georgia, 'Times New Roman', Times, serif ;
    }

    input{
        border-radius: 8px;
        height: 40px;
        width: auto;
        outline: none;
        border: none;
        color:black;
        font-size: large;
        padding-inline-start: 5%;
    }

    .label{
        display: flex;
        flex-direction: column;
    }

    .output{
        display: flex;
        flex-direction: row;
        justify-content: center;
        align-items: center;
        border: solid 2px #BDBFBF;
        border-radius: 15px;
        height: 200px;
        width: 90%;
    }

    .dataBase{
        display: flex;
        flex-direction: column;
        align-items: center;
        overflow: scroll;
        width: 40%;
        height: 35em;
        border-radius: 10px;
        border: solid white 2px;
    }

    button{
        width: 300px;
        height: 4em;
        border: none;
        border-radius: 10px;
    }
    @media screen and (max-width: 1000px){
        .input{
            border: none;
            height: 30em
        }
        .content{
            display: flex;
            flex-direction: column;
            justify-content: start;
            align-items: center;
            padding: 0;
            gap: 0;
        }
        .dataBase{
        display: grid;
        grid-template-columns: 1fr 1fr;
        grid-template-rows: 1fr 1fr 1fr 1fr 1fr;
        gap: 20px;
        height: 35em;
        width: 100%;
        border: none;
    }
    .output{
        margin-inline: 10px;
    }
    }  
</style>

<script>
    export default{
        data(){
            return{
                user:{
                    username: '',
                    idade: ''
                },
                lista: [],
            }
        },
        methods: {
            getLastValue() {
                fetch('http://127.0.0.1:8080/cadaster')
                .then(res => res.json())
                .then(body => {
                    this.lista = body
                    console.log(this.lista)
                })
            },

            getFormValues () {
                const dados = JSON.stringify(this.user)
        
                fetch('http://127.0.0.1:8080/cadaster', {
                    method: 'POST',
                    headers:{
                        'Content-Type': 'application/json'
                    },
                    body:dados
                })
            },   
        },
        mounted(){
            this.getLastValue()
        }
    }

</script>