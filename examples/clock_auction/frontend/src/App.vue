<template>
    <v-app>
        <v-content>
            <v-container fluid>
                <template v-if="isLoading">
                    Loading ...
                    <v-progress-circular
                            :size="50"
                            color="amber"
                            indeterminate
                    ></v-progress-circular>
                </template>
                <template v-else>
                    <router-view/>
                </template>
            </v-container>
        </v-content>
    </v-app>
</template>

<script>

    import router from './router'

    export default {
        name: 'App',
        router,

        data () {
            return {
                isLoading: true
            }
        },

        created () {
            this.$store.dispatch('auction/LOAD_AUCTION')
                .then(this.isLoading = false)
        }
    };
</script>
