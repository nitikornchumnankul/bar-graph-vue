

<template>
  <Bar
    :options="chartOptions"
    :data="chartData"
  />

  <p>Check value from chartdata</p>
  <p>{{ this.chartData.datasets[0].data  }}</p>
  <p>{{ this.chartData.labels  }}</p>
  
</template>


<script>
import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'
import { createApp } from 'vue'
import axios from 'axios'
ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)


export default{
  name: 'BarChart',
  components: { Bar },
  setup(){
    
    return {
      data: [],
      temp: [],
    }
  },
  data () {
      
    return {
      chartData: {
        labels: [],
        datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: []
          }
        ]
      },
    }
  },
  mounted () {
    this.fillData()
  },

  methods: {
    fillData (){
      const get_data =  axios.get('http://localhost:8080/data').then(response => {
      console.log(response.data.data)
      this.chartData = {
        labels: response.data.labels,
        datasets: [
          {
            label: 'Data One',
            backgroundColor: '#f87979',
            data: response.data.data
          }
        ]
      }
      console.log(this.chartData)
       
    })
    }
  },
}


</script>
<style scoped>
h1 {
  font-weight: 500;
  font-size: 2.6rem;
  position: relative;
  top: -10px;
}

h3 {
  font-size: 1.2rem;
}

.greetings h1,
.greetings h3 {
  text-align: center;
}

@media (min-width: 1024px) {
  .greetings h1,
  .greetings h3 {
    text-align: left;
  }
}
</style>
