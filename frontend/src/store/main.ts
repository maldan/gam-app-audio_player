import { ActionContext } from 'vuex';
import { MainTree } from '.';
import Axios from 'axios';

export interface IAudioTrack {
  name: string;
  url: string;
  size: number;
  duration: number;
}

export interface ISettings {
  volume: number;
  trackId: number;
  currentTime: number;
}

export type MainStore = {
  API_URL: string;
  list: IAudioTrack[];
  settings: ISettings;
};
export type MainActionContext = ActionContext<MainStore, MainTree>;

export default {
  namespaced: true,
  state(): MainStore {
    return {
      API_URL: process.env.VUE_APP_API_URL || `${window.location.origin}/api`,
      list: [],
      settings: {
        volume: 0.2,
        trackId: 0,
        currentTime: 0,
      },
    };
  },
  mutations: {
    SET_LIST(state: MainStore, list: IAudioTrack[]): void {
      state.list = list;
    },
    SET_SETTINGS(state: MainStore, settings: ISettings): void {
      state.settings = settings;
    },
    SET_TRACK_ID(state: MainStore, trackId: number): void {
      state.settings.trackId = trackId;
    },
    SET_VOLUME(state: MainStore, volume: number): void {
      state.settings.volume = volume;
    },
  },
  actions: {
    async getList(action: MainActionContext): Promise<void> {
      action.commit(
        'SET_LIST',
        (await Axios.get(`${action.rootState.main.API_URL}/main/list`)).data.response,
      );
    },
    async getSettings(action: MainActionContext): Promise<void> {
      action.commit(
        'SET_SETTINGS',
        (await Axios.get(`${action.rootState.main.API_URL}/main/settings`)).data.response,
      );
    },
    async saveSettings(action: MainActionContext): Promise<void> {
      await Axios.post(`${action.rootState.main.API_URL}/main/settings`, action.state.settings);
    },
    async changeTrack(action: MainActionContext, id: number): Promise<void> {
      action.commit('SET_TRACK_ID', id);
      await action.dispatch('saveSettings');
    },
    async changeVolume(action: MainActionContext, volume: number): Promise<void> {
      action.commit('SET_VOLUME', volume);
      await action.dispatch('saveSettings');
    },
  },
};
