import { useState, useEffect } from "react";
import PageTitle from "../../components/page_title";
import ImagesTable from "./images_table";
import ImageListModel from "../../models/image_list_model";
import FilterPanel from "./filter_panel";
import FilterModel from "./filter_model";
import TotalReport from "./total_report";
import IconImages from "../../components/icon_images";
import useImagesService from "@src/services/useImagesService";
import { useConfiguration } from "@src/store/configurationContext";

export default function ImagesPage() {
  const { base_url } = useConfiguration();
  const { get_images, remove_image } = useImagesService(base_url);

  const [images, setImages] = useState<ImageListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [filter, setFilter] = useState<FilterModel>({ dates: [], search_tag: "" });

  const refresh = () => {
    setLoading(true);
    get_images()
      .then((images: ImageListModel[]) => {
        setImages(images);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    console.log("page images mounted");
    refresh();
  }, []);

  const do_remove_image = (id: string) => {
    remove_image(id)
      .then(() => {
        refresh();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const on_date = (date: string) => {
    let result = [];
    if (filter.dates.includes(date)) {
      result = filter.dates.filter((d) => d !== date);
    } else {
      result = [...filter.dates, date];
    }
    setFilter({ ...filter, dates: result });
  };

  const on_filter_changed = (model: FilterModel) => {
    setFilter(model);
  };

  // make images by filter
  const images_to_render = images
    .filter((v: ImageListModel) => {
      if (filter.dates.length === 0) {
        return true;
      }
      const [date, ..._] = v.created.split(" ");
      return filter.dates.includes(date);
    })
    .filter((v: ImageListModel) => {
      if (filter.search_tag.length === 0) return true;

      for (let t of v.tags) {
        if (t.includes(filter.search_tag)) return true;
      }

      return false;
    });

  return (
    <div>
      <PageTitle onRefresh={refresh} isRefresh={loading}>
        <IconImages /> Images{" "}
      </PageTitle>
      <FilterPanel filter={filter} onChange={on_filter_changed} />
      {/* <div>
        <ButtonRefresh on_refresh={refresh} loading={loading} />
      </div> */}
      <ImagesTable images={images_to_render} on_remove={do_remove_image} on_date={on_date} />
      <TotalReport total={images_to_render.length} />
    </div>
  );
}
