/* Do not change, this code is generated from Golang structs */


export interface Group {
  id: number;
  name: string;
}
export interface Permission {
  id: number;
  resource_id: number;
  resource_parent_id: number;
  relation: number;
  user_id: number;
  group_id: number;
}