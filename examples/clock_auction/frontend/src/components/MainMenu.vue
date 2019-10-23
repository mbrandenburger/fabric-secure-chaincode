<template>
    <v-card>
        <v-toolbar
                color="cyan"
                dark
                flat
        >

            <v-toolbar-title>Participants Dashboard</v-toolbar-title>
            <template v-slot:extension>
                <v-tabs
                        v-model="tab"
                        align-with-title
                        background-color="transparent"
                >
                    <v-tabs-slider color="yellow"></v-tabs-slider>

                    <v-tab v-for="item in dashboard.tabs" :key="item">
                        {{ item.text }}
                    </v-tab>
                </v-tabs>
            </template>
        </v-toolbar>

        <v-tabs-items v-model="tab">
            <v-tab-item
                    v-for="item in dashboard.tabs"
                    :key="item"
            >
                <v-card flat>
                    <v-card-text v-text="text"></v-card-text>
                </v-card>
            </v-tab-item>
        </v-tabs-items>
    </v-card>
</template>

<script>
    export default {
        name: 'Dashboard',
        props: ['dashboard'],
    }
</script>
