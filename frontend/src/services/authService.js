import axios from "axios";

const API_URL = "http://localhost:8080";

export async function login({ email, password }) {
    const response = await axios.post(`${API_URL}/login`, {
        email,
        password,
    });
    localStorage.setItem("token", response.data.token);
    return response.data;
}

export function logout() {
    localStorage.removeItem("token");
}