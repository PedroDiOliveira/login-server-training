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
                <button class="butao" type="submit">Cadastrar</button>
            </div>
        </form>
        <div id="mensagem"></div>
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

    button{
        width: 300px;
        height: 4em;
        border: none;
        border-radius: 10px;
    }
    @media screen and (max-width: 1000px){
        .input{
            border: none;
        }
    }  
</style>

<script >
    export default{
        data(){
            return{
                user:{
                    username: '',
                    idade: ''
                },
                
            }
        },
        methods: {
            getFormValues () {
                const data = JSON.stringify(this.user)
        
                fetch('http://127.0.0.1:8080/cadaster', {
                    method: 'POST',
                    headers:{
                        'Content-Type': 'application/json'
                    },
                    body:data
                }).then(res => res.json())
                .then(data => console.log(data))
                .catch(error => console.log(error))
            
            console.log(data)
            console.log(this.user.username)
            console.log(this.user.idade)
            }
        }
    }

</script>