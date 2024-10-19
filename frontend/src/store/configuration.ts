/**
 * https://zustand-demo.pmnd.rs/
 */
// import { create } from 'zustand';

interface Configuration {
  base_url: string;
}

const default_configuration: Configuration = {
  base_url: '/',
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
