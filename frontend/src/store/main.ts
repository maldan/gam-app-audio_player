import { ActionContext } from 'vuex';
import { MainTree } from '.';
import Axios from 'axios';

export interface IAudioTrack {
  name: string;
  url: string;
}

export type MainStore = {
  API_URL: string;
  list: IAudioTrack[];
};
export type MainActionContext = ActionContext<MainStore, MainTree>;

export default {
  namespaced: true,
  state(): MainStore {
    return {
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      list: [],
    };
  },
  mutations: {
    SET_LIST(state: MainStore, list: IAudioTrack[]): void {
      state.list = list;
    },
  },
  actions: {
    async getList(action: MainActionContext): Promise<void> {
      action.commit(
        'SET_LIST',
        (await Axios.get(`${action.rootState.main.API_URL}/main/list`)).data.response,
      );
    },
  },
};
