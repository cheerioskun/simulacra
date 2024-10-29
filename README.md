# Simulacra

A modular Go package for interacting with various LLM providers, with a focus on structured outputs and tool usage.
Inspired by [Generative Agents: Interactive Simulacra of Human Behavior](https://arxiv.org/abs/2309.08508).

## Features

- Modular design for multiple LLM providers
- Built-in support for OpenRouter (OpenAI-compatible API)
- Structured output handling using instructor-go
- Simple configuration and setup

## Architecture

### Core Components

1. **World**: The main simulation environment
   - Manages agents and their interactions
   - Controls simulation flow
   - Handles plugin system

2. **Agents**: Individual actors in the simulation
   - Customizable behavior through the Agent interface
   - Memory management
   - Decision-making capabilities

3. **Plugins**: Extensible functionality
   - Agent plugins for individual agent behavior
   - World plugins for global simulation features

4. **Time Management**: Controls simulation time flow
   - Adjustable simulation speed
   - Pause/resume capabilities
   - Time conversion utilities

## Plugin System

Simulacra supports two types of plugins:

1. Agent Plugins: Extend individual agent capabilities
2. World Plugins: Add features to the simulation environment


