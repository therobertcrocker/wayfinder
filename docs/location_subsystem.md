# Location Subsystem Design Reference

## Overview
The location subsystem is designed to represent and manage the complex spatial relationships within the game world using a hierarchical graph structure. This document provides a summary of the design decisions and structure of the subsystem.

## Hierarchical Graph Structure

### 1. **Top-Level Nodes**
- Represent major areas or realms, such as the Astral Sea, Material Plane, and individual Fragments.
- Each top-level node can contain its own sub-graph.

### 2. **Sub-Graphs**
- Each top-level node, especially the Material Plane and Fragments, can have its own isolated graph.
- This graph contains nodes representing cities, points of interest, etc., and edges representing paths or connections between them.

### 3. **Portal Nodes**
- In the Astral Sea's graph, certain nodes act as "portals" or "gateways" to the Material Plane or Fragments.
- These portal nodes link to the corresponding top-level node of the Material Plane or a specific Fragment.

## Components in ECS

### 1. **Node Component**
- `NodeID`: A unique identifier for the node.
- `NodeType`: Describes the type of node (e.g., city, portal, point of interest).
- `ConnectedNodes`: A list of other nodes this node is directly connected to within the same graph.
- `LinkedNode`: For portal nodes in the Astral Sea, this attribute points to the top-level node of the Material Plane or a specific Fragment.

### 2. **Graph Component**
- `GraphID`: A unique identifier for the graph.
- `Nodes`: A list of all nodes within the graph.
- `Edges`: A list of all edges within the graph.
- `ParentNode`: If it's a sub-graph (like for a Fragment or the Material Plane), this points to the portal node in the Astral Sea that leads to it.

### 3. **Position Component**
- Stores the current node an entity is located at, represented by the `NodeID`.

### 4. **Movement Component**
- For entities that can move, this component can store information about their movement capabilities and any special movement attributes.

### 5. **Portal Component**
- For entities that are portals, this component can store information about the portal's destination node.

## Unified Coordinate System

### Structured Numerical Format
- **Astral Sea Coordinates**: Unique numerical identifier for each major node in the Astral Sea.
- **Portal Nodes**: Unique sub-identifier for each portal node within the Astral Sea.
- **Material Plane & Fragment Coordinates**: Unique numerical identifier for each major realm or fragment and a unique sub-identifier for each node within.

### Constructing the In-World Coordinate
- Given the abstract representation `(Astral Sea, PortalNode_5, Material Plane, CityNode_3)`, the in-world coordinate would be `001.5.4.003`.

### Additional Considerations
- **Fixed Length**: Use a fixed length for each segment for consistency.
- **Delimiter**: Use a delimiter (like a period `.`) to separate the different segments.
- **Translation Table**: Map the numerical identifiers to their corresponding abstract representation.

## Benefits
- **Consistency**: Standardized format for all in-world coordinates.
- **Readability**: Easily readable and distinguishable coordinates.
- **Reversibility**: Coordinates can be translated back to their abstract representation.

