<template>
    <div class="col no-pad">
        <div class="container no-pad">
            <h2 class="row fira-sans font-bold justify-content-between align-items-center">
                <div class="no-pad col-auto">Discussions</div>
                <div class="no-pad col-auto lang-select">
                    <span>LANGUAGE</span>
                    <select @change="onLanguageChanged($event.target.value)"> 
                        <option value="all">All</option>
                        <option value="cpp">C++</option>
                        <option value="java">Java</option>
                        <option value="python">Python</option>
                    </select>
                </div>
            </h2>
        </div>
        <article>TODO: fetch discussion</article>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: "DiscussionBox",
    methods: {
        onLanguageChanged(lang: string): void {
            this.$emit("langChanged", lang);
        }
    },
    created() {
        this.$emit("langChanged", "all");
    },
});
</script>

<style scoped>
h2 {
    font-size: 28px;
    border-bottom: 1px solid black;
    margin: 0px 15px;
}

.lang-select {
    font-size: 18px;
    border: 2px solid #285CBF;
    height: 23px;
    padding-right: 15px;
}

.lang-select>span {
    background-color: #285CBF;
    color: white;
    padding: 0px 15px;
    margin-right: 15px;
}

.lang-select>select {
    all: unset;
}
</style>
