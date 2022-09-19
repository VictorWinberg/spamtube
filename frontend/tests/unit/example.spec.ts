import { shallowMount } from "@vue/test-utils";
import App from "@/App.vue";

describe("HelloWorld.vue", () => {
  it("renders application name when mounted", () => {
    const wrapper = shallowMount(App);
    expect(wrapper.text()).toMatch("Spamtube");
  });
});
