import axios from 'axios';
import MockAdapter from 'axios-mock-adapter';

if(process.env.NODE_ENV === "development") {
    console.log("Axios API Mock running");
    const mock = new MockAdapter(axios);
}
