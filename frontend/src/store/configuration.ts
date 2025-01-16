/**
 * https://zustand-demo.pmnd.rs/
 */
// import { create } from 'zustand';

type Configuration = {
  base_url: string;
  version: string;
  use_registry: boolean;
};

const default_configuration: Configuration = {
  base_url: '/',
  version: '0.0.0',
  use_registry: false,
};

var store: Configuration = {
  ...default_configuration,
};

// NOTE: не работает, в компонентах при попытке прочитать вываливается ошибка компиляции
// export const useConfiguration = create((set) => ({
//   configuration: { ...default_configuration },
//   setConfiguration: (value: Configuration) => set({ configuration: value }),
// }));

// export const useStore = create((set) => ({
//   count: 1,
//   inc: () => set((state: any) => ({ count: state.count + 1 })),
// }));

export function useConfiguration() {
  const setConfiguration = (conf: Configuration) => {
    store = {
      ...default_configuration,
      ...conf,
    };
  };

  return {
    setConfiguration,
    get configuration(): Configuration {
      return { ...store };
    },
  };
}
