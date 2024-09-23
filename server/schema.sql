create table devices (
  id text not null primary key,
  version text not null,
  model text not null,
  number text not null, -- the phone number inside the device
  data_amount integer not null default 0
);

create table users (
  id text not null primary key,
  phone text not null default '',
  email text not null default '',
  password text not null
);

create table sessions (
  id text not null primary key,
  user_id text not null references users(id),
  valid boolean not null default true
);

create table subscriptions (
  id text not null primary key,
  user_id text not null references users(id),
  bought_on date not null,
  expires_on date not null
);

create table device_positions (
  device_id text not null references devices(id),
  latitude decimal(9,6) not null,
  longitude decimal(9,6) not null,
  logged_at timestamp not null default now(),
  primary key (device_id, logged_at)
);

create table admins (
  username text not null primary key,
  password text not null
);
