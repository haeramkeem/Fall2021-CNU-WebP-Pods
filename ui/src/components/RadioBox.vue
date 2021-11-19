<template>
    <ul class="d-inline-block">
        <li v-for="(item, idx) in radioItems" v-bind:key="idx" @click="onClick(idx)" class="d-inline-block">
            <div class="open-sans font-bold text-white selected" v-if="selected == idx">{{item}}</div>
            <div class="open-sans font-bold text-white" v-else>{{item}}</div>
        </li>
    </ul>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';

export default defineComponent({
    name: "RadioBox",
    props: {
        items: Array as PropType<string[]>,
    },
    data() {
        return {
            selected: 0,
        };
    },
    methods: {
        onClick(idx: number) {
            this.selected = idx;
            this.$emit("radioClicked", idx);
        }
    },
    computed: {
        radioItems() {
            if(this.items instanceof Array) {
                return this.items.map((val) => val.toUpperCase());
            }
            return [];
        }
    }
});
</script>

<style scoped>
ul {
    list-style: none;
    width: auto;
}

li>div {
    background-color: #A9A9A9;
    padding: 15px 20px;
    font-size: 20px;
    margin-left: 2px;
    cursor: pointer;
}

.selected {
    background-color: #285CBF;
}
</style>
