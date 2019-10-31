<template>
    <v-app>
        <v-content>
            <v-container fluid>
                <router-view/>
            </v-container>
        </v-content>

    </v-app>
</template>

<script>

    import router from './router'

    export default {
        name: 'App',
        router,
    };
</script>
