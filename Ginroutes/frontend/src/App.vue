<script >
import {ref}  from 'vue'
import { defineComponent } from 'vue';

import axios from 'axios';

export default defineComponent({
  data() {
    return {
      data: null,  // Puedes ajustar el tipo según la estructura de tus datos
    };
  },
  methods: {
    fetchData() {
      // Hacer una solicitud GET a una API de ejemplo
      axios.get('http://localhost:8080/data')
        .then(response => {
          // Actualizar el estado con los datos de la respuesta
          this.data = response.data;
        })
        .catch(error => {
          console.error('Error al obtener datos:', error);
        });
    },
  },
});

</script>

<template>
  <div>
    <h1>Consulta de Datos</h1>
    <button @click="fetchData">Obtener Datos</button>
    <div v-if="data">
      <h2>Respuesta del Servidor:</h2>
      <pre>{{ data }}</pre>
    </div>
  </div>
</template>

<style scoped>
  h1{
    color: blueviolet;
  }
</style>
