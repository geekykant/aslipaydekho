export default defineNuxtRouteMiddleware((to, from) => {
    if (!isAuthenticated()) {
        return navigateTo("/")
    }
});

function isAuthenticated() {
    return true;
}