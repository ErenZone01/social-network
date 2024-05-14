<template>
    <div>
      <input
        type="text"
   
        @focus="showCalendar = true"
        v-model="birthdate"
      />
      <div v-if="showCalendar" class="calendar">
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        birthdate: '',
        showCalendar: false,
        currentDate: new Date(),
        currentMonth: '',
        currentYear: '',
        days: ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']
      };
    },
    watch: {
      birthdate() {
        this.showCalendar = true; // Afficher le calendrier lorsque la date de naissance est en cours de saisie
      }
    },
    methods: {
      updateCalendar() {
        this.currentMonth = this.currentDate.toLocaleString('default', { month: 'long' });
        this.currentYear = this.currentDate.getFullYear();
      },
      previous() {
        this.currentDate.setMonth(this.currentDate.getMonth() - 1);
        this.updateCalendar();
      },
      next() {
        this.currentDate.setMonth(this.currentDate.getMonth() + 1);
        this.updateCalendar();
      }
    }
  };
  </script>
  
  <style scoped>
  </style>
  