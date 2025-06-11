import { Link } from "react-router-dom";
import { route, join_url } from "../../routes";
import { format_date } from "@src/utils/humanize";
import { Service, ServiceMode } from "@src/models/service";

type Props = {
  services: Service[];
  on_remove(id: string): void;
};

/*

example from docker
ID             NAME                                      MODE         REPLICAS   IMAGE                                                                       PORTS
mbtse0esvln5   7dbaf885_ba_w1_main_adebb59265fc4_tests   replicated   0/1        172.28.1.1:5000/kaspersky/kata/management/authorization_service/test:test   
67a1nabpe8d0   32b02965_ba_w1_main_94748aeaf2774_tests   replicated   0/1        172.28.1.1:5000/kaspersky/kata/management/authorization_service/test:test   
*/
export default function ServicesTable({ services }: Props) {
  // const on_remove_click = (e: any, name: string) => {
  //   e.preventDefault();
  //   on_remove(name);
  // };

  const rows_view = services.map((service, idx) => {
    const options_view = () => {
      return (
        <td>
          {/* TODO */}
          {/* <a href='#' className='error' onClick={(e) => on_remove_click(e, service.name)}>
            <IconRemove />
            &nbsp; Remove
          </a> */}
        </td>
      );
    };

    let image_name = service.image;
    if (image_name.length > 0) {
      const split_result = image_name.split("@");
      image_name = split_result[0];
    }

    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.service, service.id)}>{service.name}</Link>
        </td>
        <td>
          <ModeCell mode={service.mode} />
        </td>
        <td>{image_name}</td>
        {/* <td> {service.used ? 'yes' : 'no'} </td>
        <td> {service.stack_name} </td>
        <td> {service.driver} </td> */}
        {/* <!-- <td> .Mountpoint </td> --> */}
        <td> {format_date(service.created_at)} </td>
        {options_view()}
      </tr>
    );
  });
  return (
    <table className='table table-sm table-striped'>
      <thead>
        <tr>
          <th>Name</th>
          <th>Mode</th>
          {/* 
          <th>Stack</th>
          <th>Driver</th> */}
          {/* <!-- <th>Mount point</th> --> */}
          <th>Image</th>
          <th>Created</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{rows_view}</tbody>
    </table>
  );
}

type ModeCellProps = {
  mode: ServiceMode;
};
function ModeCell({ mode }: ModeCellProps) {
  if (mode.global) {
    return <span>global</span>;
  }

  if (mode.replicated) {
    // TODO: >0</1
    return <span>replicated 0/{mode.replicated.replicas}</span>;
  }

  return <span>unknown</span>;
}
