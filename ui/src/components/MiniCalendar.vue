// Show mini calendar below the header
//      props:
//          minify: Boolean
//              Size of mini calendar
//      emit:
//          event name: dateChanged
//          parameter:
//              Date: changed date object

<template>
    <div class="container-auto no-pad" id="mini-calendar">
        <div class="row align-items-center justify-content-center">
            <div v-if="this.minify" class="col-auto open-sans font-bold" id="mini-scale">{{strWeekDay}}, {{strDate}}</div>
            <div v-else class="col-auto no-pad" id="full-scale">
                <button class="arrow-btn" id="yesterday-btn" @click="gotoYesterday"></button>
                <div class="text-center" id="day-date">
                    <div class="open-sans font-reg" id="week-day">{{strWeekDay}}</div>
                    <div class="fira-sans font-bold" id="full-date">{{strDate}}</div>
                </div>
                <button class="arrow-btn" id="tomorrow-btn" @click="gotoTommorow" v-if="!isToday"></button>
                <div v-else class="arrow-btn"></div>
            </div>
        </div>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';

const date = new Date();
export default defineComponent({
    name: "MiniCalendar",
    props: {
        minify: Boolean,
    },
    data() {
        return {
            strDate: "",
            weekDay: date.getDay(),
        };
    },
    methods: {
        emitDateChanged() {
            this.$emit("dateChanged", date);
        },
        updateCal() {
            this.strDate = this.dtos(date);
            this.weekDay = date.getDay();
            this.emitDateChanged();
        },
        gotoYesterday() {
            date.setDate(date.getDate() - 1);
            this.updateCal();
        },
        gotoTommorow() {
            date.setDate(date.getDate() + 1);
            this.updateCal();
        },
        itos(digit: number) {
            if(digit < 10) {
                return "0" + digit;
            }
            return "" + digit;
        },
        dtos(target: Date) {
            return this.itos(target.getFullYear() - 2000) +
                this.itos(target.getMonth() + 1) +
                this.itos(target.getDate());
        },
    },
    computed: {
        strWeekDay() {
            const week = [
                "SUNDAY",
                "MONDAY",
                "TUESDAY",
                "WEDNESDAY",
                "THURSDAY",
                "FRIDAY",
                "SATURDAY",
            ];
            return week[this.weekDay];
        },
        isToday() {
            const today = new Date();
            return this.strDate === this.dtos(today);
        },
    },
    created() {
        this.strDate = this.dtos(date);
        this.emitDateChanged();
    },
});
</script>

<style>
#mini-calendar { width: 400px; }
#full-scale>* {
    display: inline-block;
}

#mini-scale { font-size: 24px; }

#week-day {
    color: grey;
}

#full-date {
    font-size: 60px;
}

.arrow-btn {
    width: 30px;
    height: 44px;
    border: none;
    outline: none;
    background-color: white;
    background-repeat: no-repeat;
    background-size: contain;
}

#yesterday-btn {
    background-image: url("../assets/img/left-arrow.svg");
}

#tomorrow-btn {
    background-image: url("../assets/img/right-arrow.svg");
}

#day-date {
    width: 300px;
}

.text-center * {
    text-align: center;
}
</style>
