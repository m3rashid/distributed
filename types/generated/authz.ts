/* Do not change, this code is generated from Golang structs */

export enum Relation {
  reader = 1,
  editor = 2,
  admin = 4,
  owner = 8,
}

export interface Group {
  id: number;
  name: string;
}
export interface Permission {
  id: number;
  resource_id: number;
  resource_parent_id: number;
  relation: Relation;
  user_id: number;
  group_id: number;
}
