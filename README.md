# zakame's Go Library

This started out as a practical means to learn about Go packages and
modules, but now these seem useful elsewhere:

- cache: provides a naïve in-memory cache, yet safe for concurrent
  access
- epochtime: provides an adapter for time.Time that accepts Unix epoch
  time as string, mainly for use in unmarshalling JSON to structs
