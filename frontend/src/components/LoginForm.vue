<template>
  <div class="flex min-h-screen items-center justify-center bg-gray-100">
    <div class="w-full max-w-sm rounded-lg bg-white p-6 shadow-lg">
      <h2 class="mb-4 text-center text-2xl font-bold">Login</h2>
      <form @submit.prevent="handleLogin">
        <div class="mb-4">
          <label class="mb-1 block text-sm font-medium text-gray-700">Email</label>
          <input
            v-model="email"
            type="email"
            class="w-full rounded-lg border px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>

        <div class="mb-4">
          <label class="mb-1 block text-sm font-medium text-gray-700">Senha</label>
          <input
            v-model="password"
            type="password"
            class="w-full rounded-lg border px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500"
            required
          />
        </div>

        <button
          type="submit"
          :disabled="loading"
          class="w-full rounded-lg bg-blue-600 py-2 text-white hover:bg-blue-700 disabled:opacity-50"
        >
          {{ loading ? "Entrando..." : "Entrar" }}
        </button>

        <p v-if="error" class="mt-3 text-center text-sm text-red-500">{{ error }}</p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { login } from "../services/authService";

const email = ref("");
const password = ref("");
const loading = ref(false);
const error = ref("");

const handleLogin = async () => {
  loading.value = true;
  error.value = "";

  try {
    await Login({ email: email.value, password: password.value });
    window.location.href = "index.html";
  } catch (err) {
    error.value = "Email ou senha inválidos.";
  } finally {
    loading.value = false;
  }
};
</script>
