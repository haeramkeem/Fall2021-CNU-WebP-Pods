<template>
    <div>
        <div class="container">
            <div class="row justify-content-center align-items-center main-row">
                <mini-calendar :minify="false" @dateChanged="onDateChanged"/>
            </div>
            <div class="row main-row">
                <problem-box :problem="this.problem"/>
                <discussion-box @langChanged="onLanguageChanged"/>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import td from 'typedef';
import axios from 'axios';

import MiniCalendar from '../components/MiniCalendar.vue';
import ProblemBox from '../components/ProblemBox.vue';
import DiscussionBox from '../components/DiscussionBox.vue';

export default defineComponent({
    name: "Home",
    components: {
        MiniCalendar,
        ProblemBox,
        DiscussionBox,
    },
    data() {
        return {
            problem: {} as td.ProblemStruct,
        };
    },
    methods: {
        onDateChanged(date: Date) {
            console.warn("TODO: react to date change " + date.toString());
            axios.get("/api/main/problem")
            .then((resp) => {
                const refined = resp.data as td.ResponseStruct;
                if(typeof refined.error === "string") {
                    throw new Error(refined.error);
                } else {
                    this.problem = refined.content;
                }
            }).catch((err) => {
                alert(err);
            });
        },
        onLanguageChanged(lang: string) {
            console.warn("TODO: language changed to " + lang);
        },
    },
});
</script>

<style scoped>
.main-row {
    min-height: 192px;
}
</style>
