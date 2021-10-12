import axios from 'axios';

const backend = "http://localhost:8080/";
const api = axios.create({
    baseURL: backend
});

export default {
    backend,
    api
}
