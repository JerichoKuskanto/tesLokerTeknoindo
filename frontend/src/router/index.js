import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '@/components/homePage.vue';  
import insertPage from '@/components/insertPage.vue'; 
import editPage from '@/components/editPage.vue'; 

const routes = [
  { path: '/', name: 'Home', component: HomeView },
  { path: '/insert', name: 'insertPage', component: insertPage }, 
  { path: '/edit/:id', name: 'editPage', component: editPage, props: true }, 
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
