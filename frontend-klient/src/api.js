import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

const api = axios.create({
    baseURL: API_BASE_URL,
    headers: {
        'Content-Type': 'application/json',
    },
});

export const apiService = {
    getCategories: () => api.get('/categories'),

    getMenuItems: (categoryId) => api.get(`/menu/${categoryId}`),
  
    getAllMenuItems: () => api.get('/menu'),
    
    createOrder: (orderData) => api.post('/orders', orderData),
    
    getOrders: () => api.get('/orders'),
}