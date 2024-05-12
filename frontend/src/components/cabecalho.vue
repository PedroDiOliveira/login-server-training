<template>
    <div class="header">
        <a href="#"><h1>Bem Vindo</h1></a>
        <a href="#"><i class="fa-solid fa-bars"></i></a>
    </div>    
</template>

<style scoped>
    h1{
        font-family: "Poetsen One", sans-serif;
    }

    .header a{
        color: white;
        text-decoration: none;
    }

    .header{
        display: flex;
        justify-content: space-between;
        align-items: center;
        padding-inline: 2em;
        padding-block: 2em;
    }

    i{
        font-size: xx-large;
    }
    
</style>

<script>
    export default{
        name: "cabecalho"
    }
</script>